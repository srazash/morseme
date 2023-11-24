package main

import (
	"fmt"
	"morseme/server/ticket"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.Static("/", "static")

	e.GET("/ticket", func(c echo.Context) error {
		markup := fmt.Sprintf("<pre id=\"ticket-no\" class=\"big\">%s</pre>", ticket.GenerateTicketNo())
		return c.HTML(http.StatusOK, markup)
	})

	e.Logger.Fatal(e.Start(":3000"))
}
