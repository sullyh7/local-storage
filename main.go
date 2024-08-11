package main

import (
	"embed"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sullyh7/local-storage/handler"
)

//go:embed public
var FS embed.FS

func main() {
	e := echo.New()
	e.Debug = true
	e.Pre(middleware.RemoveTrailingSlash())
	e.StaticFS("/", FS)

	e.GET("/", handler.HandleHomeIndex)

	e.Start(":3000")
}
