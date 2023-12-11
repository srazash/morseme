package main

import (
	"bytes"
	"context"
	"fmt"
	"morseme/server/morsecode"
	"morseme/server/templates"
	"net/http"
	"time"

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
		return c.HTML(http.StatusOK, fmt.Sprintf(`<footer
		hx-get="/footer">
		%s %d
		</footer>`, "Ryan Shaw-Harrison,", time.Now().Year()))
	})

	e.GET("/title-morse", func(c echo.Context) error {
		return c.HTML(http.StatusOK, `<h1 id="h1-title" hx-get="/title-text" hx-trigger="click" hx-target="#h1-title" hx-swap="outerHTML">-- --- .-. ... . -- .</h1>`)
	})

	e.GET("/title-text", func(c echo.Context) error {
		return c.HTML(http.StatusOK, `<h1 id="h1-title" hx-get="/title-morse" hx-trigger="click" hx-target="#h1-title" hx-swap="outerHTML">MorseMe</h1>`)
	})

	e.POST("/encode-to-morse", func(c echo.Context) error {
		pre := ""
		enc, err := morsecode.Encode(c.FormValue("text-input"))
		if err != nil {
			pre = `<pre id="encode-output" class="big">Please only use letters and spaces!</pre>`
		} else {
			pre = fmt.Sprintf(`<pre id="encode-output" class="big">%s</pre>`, enc)
		}

		m := fmt.Sprintf(`<figure id="encode-figure">
		<h2>Text to Morse Test</h2>
		<form id="encode-form" class="box rows"
			hx-target="#encode-figure"
			hx-swap="outerHTML"
			hx-post="/encode-to-morse">
				<label for="text-input">
					Input:
					<input id="text-input" name="text-input" type="text" size="40" spellcheck="true" placeholder="letters and spaces only" maxlength="100" required />
				</label>
			<input class="big" type="submit" value="Encode!" />
		</form>
		%s
	</figure>`, pre)

		return c.HTML(http.StatusOK, m)
	})

	e.Logger.Fatal(e.Start(":3000"))
}
