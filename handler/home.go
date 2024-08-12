package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/sullyh7/local-storage/config"
	"github.com/sullyh7/local-storage/service/discord"
	"github.com/sullyh7/local-storage/view/home"
)

func HandleHomeIndex(c echo.Context) error {
	return Render(c, http.StatusOK, home.Index())
}

func HandleUploadAction(c echo.Context) error {
	log.Info("uploading new file")
	name := c.FormValue("name")
	f, err := c.FormFile("file")
	splitted := strings.Split(f.Filename, ".")
	ext := splitted[len(splitted)-1]
	name = name + "." + ext
	log.Infof("uploading new file with name %s and file %s\n", name, f.Filename)
	if err != nil {
		return err
	}
	file, err := f.Open()
	if err != nil {
		return err
	}
	defer file.Close()

	if err := discord.UploadFile(config.Config.DiscordChannelID, name, file); err != nil {
		log.Errorf("error uploading file: %s", err.Error())
		return Render(c, http.StatusInternalServerError, home.Form(home.FormProps{Message: "error uploading file", Success: false}))
	}
	return Render(c, http.StatusOK, home.Form(home.FormProps{Message: "uploaded", Success: true}))
}

func HandleGetFiles(c echo.Context) error {
	ms, err := discord.GetFiles(config.Config.DiscordChannelID)
	if err != nil {
		return err
	}

	return Render(c, http.StatusOK, home.Files(ms))
}

func HandleDeleteAction(c echo.Context) error {
	fmt.Println("deleting...")
	id := c.Param("id")
	log.Infof("deleting with is %d", id)

	if err := discord.DeleteFile(config.Config.DiscordChannelID, id); err != nil {
		return err
	}
	return nil
}
