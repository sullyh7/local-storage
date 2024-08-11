package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sullyh7/local-storage/view/home"
)

func HandleHomeIndex(c echo.Context) error {
	return Render(c, http.StatusOK, home.Index())
}
