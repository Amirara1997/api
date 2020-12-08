package handler



import (
"github.com/labstack/echo"


)

var status_201 = response_json{
	"code":    "201",
	"message": "data_created",
}
var status_304 = response_json{
	"code":    "304",
	"message": "not_modified",
}
var status_401 = response_json{
	"code":    "401",
	"message": "unauthorized",
}
var status_404 = response_json{
	"code":    "404",
	"message": "server_not_found",
}
var status_200 = response_json{
	"code":    "200",
	"message": "OK",
}
var status_500 = response_json{
	"code":    "500",
	"message": "internal_server_error",
}



func intertnal_server_error(c echo.Context) error {
	response := response_json{
		"status": status_500,
	}
	return c.JSON(200, response)
}

func server_not_found(c echo.Context) error {
	response := response_json{
		"status": status_404,
	}
	return c.JSON(200, response)
}

func not_modified(c echo.Context) error {
	response := response_json{
		"status": status_304,
	}
	return c.JSON(200, response)

}
