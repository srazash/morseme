package main

import (
	"bytes"
	"context"
	"fmt"
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

	e.GET("/check", func(c echo.Context) error {
		m := new(bytes.Buffer)
		templates.NewCheck().Render(context.Background(), m)
		return c.HTML(http.StatusOK, m.String())
	})

	e.GET("/stats", func(c echo.Context) error {
		m := new(bytes.Buffer)
		templates.MessageStats(message.MessageStats()).Render(context.Background(), m)
		return c.HTML(http.StatusOK, m.String())
	})

	e.POST("/check-message", func(c echo.Context) error {
		m := new(bytes.Buffer)
		templates.GetCheck(message.CheckIMS(c.FormValue("ticket-number"))).Render(context.Background(), m)
		return c.HTML(http.StatusOK, m.String())
	})

	e.POST("/submit-message", func(c echo.Context) error {
		m := new(bytes.Buffer)

		msg, err := message.MessageHandler(c.FormValue("message-body"), c.FormValue("message-sender"))
		if err != nil {
			templates.ErrorMessage().Render(context.Background(), m)
		} else {
			c.Response().Header().Add("HX-Trigger", "SubmitMessage")
			templates.SubmitMessage(msg).Render(context.Background(), m)
		}
		return c.HTML(http.StatusOK, m.String())
	})

	// APIs
	e.GET("/api/messages/count", func(c echo.Context) error {
		m, n, o := message.MessageStats()
		return c.String(http.StatusOK, fmt.Sprintf("t:%d\tu:%d\td:%d\n", m, n, o))
	})

	e.GET("/api/messages/total", func(c echo.Context) error {
		m, _, _ := message.MessageStats()
		return c.String(http.StatusOK, fmt.Sprintf("%d\n", m))
	})

	e.GET("/api/messages/undelivered", func(c echo.Context) error {
		_, m, _ := message.MessageStats()
		return c.String(http.StatusOK, fmt.Sprintf("%d\n", m))
	})

	e.GET("/api/messages/delivered", func(c echo.Context) error {
		_, _, m := message.MessageStats()
		return c.String(http.StatusOK, fmt.Sprintf("%d\n", m))
	})

	e.Logger.Fatal(e.Start(":3000"))
}
