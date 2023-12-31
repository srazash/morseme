package main

import (
	"bytes"
	"context"
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

	e.Static("/", "static")

	e.GET("/ticket", func(c echo.Context) error {
		m := new(bytes.Buffer)
		templates.TicketNo().Render(context.Background(), m)
		return c.HTML(http.StatusOK, m.String())
	})

	e.GET("/footer", func(c echo.Context) error {
		m := new(bytes.Buffer)
		templates.Footer().Render(context.Background(), m)
		return c.HTML(http.StatusOK, m.String())
	})

	e.GET("/title-morse", func(c echo.Context) error {
		m := new(bytes.Buffer)
		templates.TitleMorse().Render(context.Background(), m)
		return c.HTML(http.StatusOK, m.String())
	})

	e.GET("/title-text", func(c echo.Context) error {
		m := new(bytes.Buffer)
		templates.TitleText().Render(context.Background(), m)
		return c.HTML(http.StatusOK, m.String())
	})

	e.POST("/encode-to-morse", func(c echo.Context) error {
		m := new(bytes.Buffer)
		enc, err := morsecode.Encode(c.FormValue("text-input"))
		templates.Encode(enc, err).Render(context.Background(), m)
		return c.HTML(http.StatusOK, m.String())
	})

	e.GET("/new-message", func(c echo.Context) error {
		m := new(bytes.Buffer)
		templates.NewMessage().Render(context.Background(), m)
		return c.HTML(http.StatusOK, m.String())
	})

	e.GET("/new-check", func(c echo.Context) error {
		m := new(bytes.Buffer)
		templates.NewCheck().Render(context.Background(), m)
		return c.HTML(http.StatusOK, m.String())
	})

	e.POST("/message", func(c echo.Context) error {
		m := message.MessageHandler(c.FormValue("message-body"), c.FormValue("message-sender"))
		return c.HTML(http.StatusOK, m)
	})

	e.POST("/check", func(c echo.Context) error {
		m := message.StringifyMessage(message.CheckIMS(c.FormValue("ticket-number")))
		return c.HTML(http.StatusOK, m)
	})

	e.GET("/stats", func(c echo.Context) error {
		m := new(bytes.Buffer)
		templates.MessageStats(message.MessageStats()).Render(context.Background(), m)
		return c.HTML(http.StatusOK, m.String())
	})

	e.Logger.Fatal(e.Start(":3000"))
}
