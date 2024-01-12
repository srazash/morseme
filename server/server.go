package main

import (
	"bytes"
	"context"
	"morseme/server/api"
	"morseme/server/message"
	"morseme/server/morsecode"
	"morseme/server/templates"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

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
		templates.MessageStats(message.MessageStats()).Render(context.Background(), html)
		return c.HTML(http.StatusOK, html.String())
	})

	e.POST("/check-message", func(c echo.Context) error {
		html := new(bytes.Buffer)
		templates.GetCheck(message.CheckIMS(c.FormValue("ticket-number"))).Render(context.Background(), html)
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
		}
		return c.HTML(http.StatusOK, html.String())
	})

	// APIs
	e.GET("/api/messages", func(c echo.Context) error {
		j := api.MessagesJson(message.MessageStore)
		return c.JSONBlob(http.StatusOK, j)
	})

	e.GET("/api/messages/latest", func(c echo.Context) error {
		j := api.LastMessageJson(message.MessageStore)
		return c.JSONBlob(http.StatusOK, j)
	})

	e.GET("/api/messages/nexttodeliver", func(c echo.Context) error {
		j := api.FirstUndeliveredMessageJson(message.MessageStore)
		return c.JSONBlob(http.StatusOK, j)
	})

	e.GET("/api/stats", func(c echo.Context) error {
		t, u, d := message.MessageStats()
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

	PopulateIms() // insert dummy messages

	e.Logger.Fatal(e.Start(":3000"))
}

func PopulateIms() {
	message.MessageHandler("hello", "rsh")
	message.MessageStore[0].DeliveredState = true
	message.MessageHandler("hi", "ryan")
	message.MessageHandler("hej", "rybear")
}
