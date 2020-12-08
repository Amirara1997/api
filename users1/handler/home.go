package handler

import (
	"github.com/labstack/echo"
	"net/http"
)

type response_json map[string]interface{}
type DB struct {
	Value       interface{}
	Error       error
	RowAffected int64
}


func Home(c echo.Context) error  {
	return c.JSON(http.StatusOK,"hello")
}