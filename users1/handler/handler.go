package handler

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"net/http"
	"users1/database"
	"users1/model"
)

func GetUsers(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var u []*model.User

		if err := db.Find(&u).Error; err != nil {
			return err
		}

		return c.JSON(http.StatusOK, u)
	}
}




func AddUserController(c echo.Context) error  {
	db,_ := database.NewDB()
	defer db.Close()

	name := c.FormValue("name")
	passwd := c.FormValue("passwd")

	user := model.User{
		Name: name,
		Passwd: passwd,
	}
	db.NewRecord(user)

	if error_insert := db.Create(&user);error_insert.Error != nil {
		return c.JSON(http.StatusBadRequest,echo.Map{"error":error_insert})
	}
	db.NewRecord(user)

	response := response_json{
		"status" : status_200,
	}
	return c.JSON(200,response)
}



func EditeDataUsersController(c echo.Context) error  {
	db,_ := database.NewDB()
	defer db.Close()

	id := c.Param("id")
	username := c.FormValue("username")
	passwd := c.FormValue("passwd")

	var user model.User

	data := db.Model(&user).Where("id = ?",id).Updates(map[string]interface{}{
		"usename": username,
		"passwd":passwd,
	})
	if data.Error != nil {
		return c.JSON(http.StatusBadRequest,echo.Map{"error":data.Error})
	} else if data.RowsAffected == 0 {
		not_modified(c)
		return nil
	}

	response := response_json{
		"status": status_200,
	}
	return c.JSON(200,response)

}