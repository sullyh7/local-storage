package main

import (
	"embed"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sullyh7/local-storage/config"
	"github.com/sullyh7/local-storage/handler"
	"github.com/sullyh7/local-storage/service/discord"
)

//go:embed public
var FS embed.FS

func main() {
	if err := config.LoadConfig(); err != nil {
		log.Fatal(err)
	}

	if err := discord.InitSession(config.Config); err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	e.Debug = true
	e.Pre(middleware.RemoveTrailingSlash())
	e.StaticFS("/", FS)

	e.GET("/", handler.HandleHomeIndex)
	e.POST("/action/upload", handler.HandleUploadAction)
	e.GET("/action/get-files", handler.HandleGetFiles)
	e.DELETE("/action/delete/:id", handler.HandleDeleteAction)

	e.Start(":3000")
}
