package main

import (
	"bytes"
	"context"
	"morseme/server/api"
	"morseme/server/api/restricted"
	"morseme/server/db"
	"morseme/server/message"
	"morseme/server/morsecode"
	"morseme/server/templates"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	db.InitDb()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Gzip())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))

	e.Static("/", "public")

	e.GET("/ticket", func(c echo.Context) error {
		html := new(bytes.Buffer)
		templates.TicketNo().Render(context.Background(), html)
		return c.HTML(http.StatusOK, html.String())
	})

	e.GET("/footer", func(c echo.Context) error {
		html := new(bytes.Buffer)
		templates.Footer().Render(context.Background(), html)
		return c.HTML(http.StatusOK, html.String())
	})

	e.GET("/title-morse", func(c echo.Context) error {
		html := new(bytes.Buffer)
		templates.TitleMorse().Render(context.Background(), html)
		return c.HTML(http.StatusOK, html.String())
	})

	e.GET("/title-text", func(c echo.Context) error {
		html := new(bytes.Buffer)
		templates.TitleText().Render(context.Background(), html)
		return c.HTML(http.StatusOK, html.String())
	})

	e.POST("/encode-to-morse", func(c echo.Context) error {
		html := new(bytes.Buffer)
		enc, err := morsecode.Encode(c.FormValue("text-input"))
		templates.Encode(enc, err).Render(context.Background(), html)
		return c.HTML(http.StatusOK, html.String())
	})

	e.GET("/new-message", func(c echo.Context) error {
		html := new(bytes.Buffer)
		templates.NewMessage().Render(context.Background(), html)
		return c.HTML(http.StatusOK, html.String())
	})

	e.GET("/check", func(c echo.Context) error {
		html := new(bytes.Buffer)
		templates.NewCheck().Render(context.Background(), html)
		return c.HTML(http.StatusOK, html.String())
	})

	e.GET("/stats", func(c echo.Context) error {
		html := new(bytes.Buffer)
		templates.MessageStats(db.MessageCount()).Render(context.Background(), html)
		return c.HTML(http.StatusOK, html.String())
	})

	e.POST("/check-message", func(c echo.Context) error {
		html := new(bytes.Buffer)
		msg, err := db.CheckMessage(c.FormValue("ticket-number"))
		templates.GetCheck(msg, err).Render(context.Background(), html)
		return c.HTML(http.StatusOK, html.String())
	})

	e.POST("/submit-message", func(c echo.Context) error {
		html := new(bytes.Buffer)

		msg, err := message.MessageHandler(c.FormValue("message-body"), c.FormValue("message-sender"))
		if err != nil {
			templates.ErrorMessage().Render(context.Background(), html)
		} else {
			c.Response().Header().Add("HX-Trigger", "SubmitMessage")
			templates.SubmitMessage(msg).Render(context.Background(), html)
			db.UpdateMessageCount()
		}
		return c.HTML(http.StatusOK, html.String())
	})

	// APIs
	e.GET("/api/stats", func(c echo.Context) error {
		t, u, d := db.MessageCount()
		j := api.MessageStatsJson(t, u, d)
		return c.JSONBlob(http.StatusOK, j)
	})

	e.GET("/api/stats/total", func(c echo.Context) error {
		t := message.MessageStatsTotal()
		j := api.MessageStatsTotalJson(t)
		return c.JSONBlob(http.StatusOK, j)
	})

	e.GET("/api/stats/undelivered", func(c echo.Context) error {
		u := message.MessageStatsUndelivered()
		j := api.MessageStatsUndeliveredJson(u)
		return c.JSONBlob(http.StatusOK, j)
	})

	e.GET("/api/stats/delivered", func(c echo.Context) error {
		d := message.MessageStatsDelivered()
		j := api.MessageStatsDeliveredJson(d)
		return c.JSONBlob(http.StatusOK, j)
	})

	// Restricted API
	e.POST("/api/login", func(c echo.Context) error {
		username := c.FormValue("username")
		password := c.FormValue("password")

		users := restricted.LoadUsers()
		pwd := users[username]

		if pwd == "" || password != pwd {
			return echo.ErrUnauthorized
		}

		claims := &restricted.JwtCustomClaims{
			Name:  username,
			Admin: true,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 5)),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		t, err := token.SignedString([]byte(restricted.SIGNING_KEY_SECRET))
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, echo.Map{
			"token": t,
		})
	})

	r := e.Group("/restricted")
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(restricted.JwtCustomClaims)
		},
		SigningKey: []byte(restricted.SIGNING_KEY_SECRET),
	}
	r.Use(echojwt.WithConfig(config))

	r.GET("/api/messages", func(c echo.Context) error {
		j := api.MessagesJson(message.MessageStore)
		return c.JSONBlob(http.StatusOK, j)
	})

	r.GET("/api/messages/latest", func(c echo.Context) error {
		j := api.LastMessageJson(message.MessageStore)
		return c.JSONBlob(http.StatusOK, j)
	})

	r.GET("/api/messages/nexttodeliver", func(c echo.Context) error {
		j := api.FirstUndeliveredMessageJson(message.MessageStore)
		return c.JSONBlob(http.StatusOK, j)
	})

	e.Logger.Fatal(e.Start(":3000"))
}
