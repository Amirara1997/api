package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"html/template"
	"io"
	"log"
	"github.com/foolin/echo-template"
	"net/http"
	"users1/database"
	"users1/handler"
)

type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}



func main()  {
	db, err := database.NewDB()
	logFatal(err)

	db.LogMode(true)
	defer db.Close()

	e := echo.New()

	e.Renderer = echotemplate.Default()
	e.GET("/page", func(c echo.Context) error {
		//render only file, must full name with extension
		return c.Render(http.StatusOK, "home.html", echo.Map{"title": "Page file title!!"})
	})
	//e.GET("/", handler.Home)
	e.GET("/user", handler.GetUsers(db))
	e.POST("/user",handler.AddUserController)
	e.POST("/user/:id",handler.EditeDataUsersController)
	e.POST("/login", handler.Login)
	e.GET("/logout", handler.Logout)


	admingroup := e.Group("/admin")
	admingroup.Use(middleware.BasicAuth(func(username, passwd string, c echo.Context) (bool, error) {
		if username == "admin" && passwd == "admin" {
			return true,nil
		}
		return true,nil
	}))
	admingroup.GET("", handler.Admin)

	err = e.Start(":8000")
	logFatal(err)
}

func logFatal(err error)  {
	if err != nil {
		log.Fatalln(err)
	}
}
