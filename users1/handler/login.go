package handler



import (


	"net/http"
	"os"
	"time"
	"users1/model"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

var err error

func Login(c echo.Context) error {
	name := c.FormValue("name")
	passwd := c.FormValue("passwd")

	user, err := model.GetUserByUsername(name)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Passwd), []byte(passwd))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{
			"message": "Invalid username / password",
		})
	}

	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd")
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["id"] = user.ID
	claims["username"] = user.Name
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))

	if err != nil {
		return err
	}

	c.SetCookie(&http.Cookie{
		Name:     "Authorization",
		Value:    token,
		Path:     "/",
		Expires:  time.Now().Add(24 * time.Hour),
		Secure:   true,
		HttpOnly: true,
	})

	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}
