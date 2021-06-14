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

func (uc UserController) ShowUserProfile(c echo.Context) error {
	user, ok := c.Get("user").(*entity.User)
	if !ok {
		log.Println("Failed to parse user session")
		return echo.ErrUnauthorized
	}

	data := echo.Map{"User": user}
	return renderPage(c, "profile", user.Name, data)
}

func (uc UserController) ShowEditUserPasswordForm(c echo.Context) error {
	return renderPage(c, "profile_password", "Change Password", nil)
}

func (uc UserController) ShowEditUserProfileForm(c echo.Context) error {
	user, ok := c.Get("user").(*entity.User)
	if !ok {
		log.Println("Failed to parse user session")
		return echo.ErrUnauthorized
	}

	data := echo.Map{"User": user}
	return renderPage(c, "profile_edit", "Edit Profile", data)
}

func (uc UserController) UpdateUserProfile(c echo.Context) error {
	// photo, err := c.FormFile("photo")
	// if err != nil {
	// 	return echo.ErrBadRequest
	// }

	return nil
}

func (uc UserController) UpdateUserPassword(c echo.Context) error {
	sess, _ := session.Get("GO-POS", c)

	var updatePasswordParam entity.UpdateUserPasswordParam
	if err := c.Bind(&updatePasswordParam); err != nil {
		return echo.ErrInternalServerError
	}

	err := c.Validate(&updatePasswordParam)
	if ev, ok := err.(entity.ErrValidation); ok {
		sess.AddFlash(ev, "error_validation")
		if err := sess.Save(c.Request(), c.Response()); err != nil {
			log.Println(err)
		}
		return c.Redirect(http.StatusSeeOther, "/profile/edit/password")
	}

	if err != nil {
		return echo.ErrInternalServerError
	}

	user, ok := c.Get("user").(*entity.User)
	if !ok {
		log.Println("Failed to parse user session")
		return echo.ErrUnauthorized
	}

	isUpdated, err := uc.userUsecase.UpdateUserPassword(user.ID, updatePasswordParam.Password)
	if err != nil {
		return err
	}

	if !isUpdated {
		return echo.ErrInternalServerError
	}

	sess.AddFlash("Success updating password", "success_message")
	sess.Save(c.Request(), c.Response())
	return c.Redirect(http.StatusSeeOther, "/profile")
}

func (uc UserController) Login(c echo.Context) error {
	sess, _ := session.Get("GO-POS", c)

	var credential entity.UserCredential
	if err := c.Bind(&credential); err != nil {
		return echo.ErrInternalServerError
	}

	err := c.Validate(&credential)
	if ev, ok := err.(entity.ErrValidation); ok {
		sess.AddFlash(ev, "error_validation")
		if err := sess.Save(c.Request(), c.Response()); err != nil {
			log.Println(err)
		}
		return c.Redirect(http.StatusSeeOther, "/auth/login")
	}

	if err != nil {
		return echo.ErrInternalServerError
	}

	user, err := uc.userUsecase.GetUserByCredential(credential)
	_, isent := err.(entity.ErrNotFound)
	_, iseic := err.(entity.ErrInvalidCredential)
	if isent || iseic {
		sess.AddFlash("Incorrect Email or Password", "error_message")
		sess.Save(c.Request(), c.Response())
		return c.Redirect(http.StatusSeeOther, "/auth/login")
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
