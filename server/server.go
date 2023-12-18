package main

import (
	"bytes"
	"context"
	"fmt"
	"morseme/server/morsecode"
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
		d := ""
		enc, err := morsecode.Encode(c.FormValue("text-input"))
		if err != nil {
			d = `<div id="encode-output" class="terminal-alert terminal-alert-error">invalid input: letters and spaces only!</div>`
		} else {
			d = fmt.Sprintf(`<div id="encode-output" class="terminal-alert terminal-alert-primary">%s</div>`, enc)
		}

		m := fmt.Sprintf(`<form id="encode-form"
		hx-target="#encode-form"
		hx-swap="outerHTML"
		hx-post="/encode-to-morse">
		<fieldset>
			<legend>Text to convert</legend>
			<div>
			<label for="text-input">
				Input:
				<input id="text-input" name="text-input" type="text" size="40" spellcheck="true" placeholder="letters and spaces only" maxlength="100" required />
			</label>
			</div>
			<br>
			<div>
				<input id="submit-btn" class="btn btn-default" type="submit" value="Convert" />
			</div>
		</fieldset>
		<br>
		%s
		</form>`, d)

		return c.HTML(http.StatusOK, m)
	})

	e.Logger.Fatal(e.Start(":3000"))
}
