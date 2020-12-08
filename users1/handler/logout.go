package handler


import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"

	"net/http"
)

func Logout(c echo.Context) error  {
	sess ,_ := session.Get("Authorization",c)
	sess.Options = &sessions.Options{
		Path: "/",
		MaxAge: -1,

		HttpOnly: true,

	}
	sess.Values["name"]= ""
	sess.Values["authenticated"]= ""
	sess.Save(c.Request(),c.Response())
	return c.Redirect(http.StatusSeeOther,"/")


}