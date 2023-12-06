package main

import (
	"bytes"
	"context"
	"morseme/server/templates"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.Static("/", "static")

	e.GET("/ticket", func(c echo.Context) error {
		m := new(bytes.Buffer)
		templates.TicketNo().Render(context.Background(), m)
		return c.HTML(http.StatusOK, m.String())
	})

	e.GET("/title-morse", func(c echo.Context) error {
		return c.HTML(http.StatusOK, `<h1 id="h1-title" hx-get="/title-text" hx-trigger="click" hx-target="#h1-title" hx-swap="outerHTML">-- --- .-. ... . -- .</h1>`)
	})

	e.GET("/title-text", func(c echo.Context) error {
		return c.HTML(http.StatusOK, `<h1 id="h1-title" hx-get="/title-morse" hx-trigger="click" hx-target="#h1-title" hx-swap="outerHTML">MorseMe</h1>`)
	})

	e.POST("/encode-to-morse", func(c echo.Context) error {
		return c.HTML(http.StatusOK, `<pre id="encode-output" class="big">Submitted!</pre>`)
	})

	e.Logger.Fatal(e.Start(":3000"))
}
