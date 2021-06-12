package controller

import (
	"log"
	"net/http"

	"github.com/ardafirdausr/go-pos/internal"
	"github.com/ardafirdausr/go-pos/internal/app"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	userUsecase internal.UserUsecase
}

func NewUserController(app *app.Usecases) *UserController {
	return &UserController{userUsecase: app.UserUsecase}
}

func (uc UserController) ShowLoginForm(c echo.Context) error {
	return renderPage(c, "login", "Login", nil)
}

func (uc UserController) Login(c echo.Context) error {
	email := c.Request().FormValue("email")
	password := c.Request().FormValue("password")

	user, err := uc.userUsecase.GetUserByCredential(email, password)
	if err != nil {
		log.Println(err)
		return err
	}

	sess, err := session.Get("GO-POS", c)
	if err != nil {
		log.Println(err)
		return err
	}

	if sess.IsNew {
		sess.Options = &sessions.Options{
			Path:     "/",
			MaxAge:   86400 * 7,
			HttpOnly: true,
		}
	}

	sess.Values["user"] = user
	sess.AddFlash("message", "Login Success")
	if err := sess.Save(c.Request(), c.Response()); err != nil {
		log.Println(err.Error())
		return err
	}
	return c.Redirect(http.StatusSeeOther, "/dashboard")
}

func (uc UserController) Logout(c echo.Context) error {
	sess, err := session.Get("GO-POS", c)
	if err != nil {
		log.Println(err)
	}

	sess.Options.MaxAge = -1
	if err := sess.Save(c.Request(), c.Response()); err != nil {
		log.Println(err)
	}

	return c.Redirect(http.StatusSeeOther, "/")
}
