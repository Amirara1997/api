package handler

import (
	"net/http"
	"github.com/labstack/echo"
)

func Admin(c echo.Context) error  {
	return c.JSON(http.StatusOK, "admin page")
}