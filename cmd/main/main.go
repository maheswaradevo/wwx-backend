package main

import (
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(ctx echo.Context) error {
		return ctx.String(200, "Hello World")
	})
	e.Logger.Fatal(e.Start(":8080"))
}
