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
	db.LoadUsersToDb()

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
		templates.MessageStats(db.ReadMessageCountCache()).Render(context.Background(), html)
		return c.HTML(http.StatusOK, html.String())
	})

	e.POST("/check-message", func(c echo.Context) error {
		html := new(bytes.Buffer)
		msg, err := db.ReadMessage(c.FormValue("ticket-number"))
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
			db.UpdateMessageCountCache()
		}
		return c.HTML(http.StatusOK, html.String())
	})

	// APIs
	e.GET("/api/stats", func(c echo.Context) error {
		t, u, d := db.ReadMessageCountCache()
		j := api.MessageStatsJson(t, u, d)
		return c.JSONBlob(http.StatusOK, j)
	})

	e.GET("/api/stats/total", func(c echo.Context) error {
		t, _, _ := db.ReadMessageCountCache()
		j := api.MessageStatsTotalJson(t)
		return c.JSONBlob(http.StatusOK, j)
	})

	e.GET("/api/stats/undelivered", func(c echo.Context) error {
		_, u, _ := db.ReadMessageCountCache()
		j := api.MessageStatsUndeliveredJson(u)
		return c.JSONBlob(http.StatusOK, j)
	})

	e.GET("/api/stats/delivered", func(c echo.Context) error {
		_, _, d := db.ReadMessageCountCache()
		j := api.MessageStatsDeliveredJson(d)
		return c.JSONBlob(http.StatusOK, j)
	})

	// Restricted API
	e.POST("/api/login", func(c echo.Context) error {
		username := c.FormValue("username")
		password := c.FormValue("password")

		users := db.ReadAllUsersMap()
		pwd := users[username]

		if pwd == "" || password != pwd {
			return echo.ErrUnauthorized
		}

		claims := &restricted.JwtCustomClaims{
			Name:  username,
			Admin: true,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 15)),
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

	r := e.Group("/res")
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(restricted.JwtCustomClaims)
		},
		SigningKey: []byte(restricted.SIGNING_KEY_SECRET),
	}
	r.Use(echojwt.WithConfig(config))

	r.GET("/messages", func(c echo.Context) error {
		m := db.ReadAllMessages()
		j := api.MessagesJson(m)
		return c.JSONBlob(http.StatusOK, j)
	})

	r.GET("/latest", func(c echo.Context) error {
		m, _ := db.ReadLatestMessage()
		j := api.MessageJson(m)
		return c.JSONBlob(http.StatusOK, j)
	})

	r.GET("/next", func(c echo.Context) error {
		m, _ := db.ReadFirstUndeliveredMessage()
		j := api.MessageJson(m)
		return c.JSONBlob(http.StatusOK, j)
	})

	r.GET("/deliver", func(c echo.Context) error {
		m, _ := db.DeliverMessage()
		j := api.MessageJson(m)
		return c.JSONBlob(http.StatusOK, j)
	})

	e.Logger.Fatal(e.Start(":3000"))
}
