package controller

import (
	"log"
	"net/http"

	"github.com/ardafirdausr/go-pos/internal"
	"github.com/ardafirdausr/go-pos/internal/app"
	"github.com/ardafirdausr/go-pos/internal/entity"
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
	sess, _ := session.Get("GO-POS", c)

	email := c.Request().FormValue("email")
	password := c.Request().FormValue("password")

	user, err := uc.userUsecase.GetUserByCredential(email, password)
	if _, ok := err.(entity.ErrNotFound); ok {
		sess.AddFlash("Incorrect Email or Password", "error_message")
		sess.Save(c.Request(), c.Response())
		c.Redirect(http.StatusSeeOther, "/auth/login")
	}

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
