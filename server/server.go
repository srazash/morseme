package main

import (
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.Static("/", "static")

	e.Logger.Fatal(e.Start(":3000"))
}
