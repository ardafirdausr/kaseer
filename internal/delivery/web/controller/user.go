package controller

import (
	"log"
	"net/http"

	"github.com/ardafirdausr/kaseer/internal"
	"github.com/ardafirdausr/kaseer/internal/app"
	"github.com/ardafirdausr/kaseer/internal/entity"
	"github.com/ardafirdausr/kaseer/internal/pkg/storage"
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
	return renderPage(c, "profile", "Profile", nil)
}

func (uc UserController) ShowEditUserPasswordForm(c echo.Context) error {
	return renderPage(c, "profile_password", "Change Password", nil)
}

func (uc UserController) ShowEditUserProfileForm(c echo.Context) error {
	return renderPage(c, "profile_edit", "Edit Profile", nil)
}

func (uc UserController) UpdateUserProfile(c echo.Context) error {
	sess, _ := session.Get("kaseer", c)

	user, ok := c.Get("user").(*entity.User)
	if !ok {
		return echo.ErrInternalServerError
	}

	if user.Email == "staff@mail.com" {
		sess.AddFlash("Cannot update this demo account", "error_message")
		sess.Save(c.Request(), c.Response())
		return c.Redirect(http.StatusSeeOther, "/profile/edit")
	}

	if err := c.Bind(user); err != nil {
		return echo.ErrInternalServerError
	}

	param := entity.UpdateUserParam{
		Name:     user.Name,
		Email:    user.Email,
		PhotoUrl: user.PhotoUrl,
	}

	photo, _ := c.FormFile("photo")
	if photo != nil {
		strg := storage.FileSystemStorage{}
		photoUrl, err := uc.userUsecase.SaveUserPhoto(strg, user, photo)
		if err != nil {
			sess.AddFlash("Failed upload photo", "error_message")
			return c.Redirect(http.StatusSeeOther, "/profile/edit")
		}

		param.PhotoUrl = &photoUrl
	}

	err := c.Validate(&param)
	if ev, ok := err.(entity.ErrValidation); ok {
		sess.AddFlash(ev, "error_validation")
		if err := sess.Save(c.Request(), c.Response()); err != nil {
			log.Println(err)
		}
		return c.Redirect(http.StatusSeeOther, "/profile/edit")
	}

	if err != nil {
		return echo.ErrInternalServerError
	}

	isUpdated, err := uc.userUsecase.UpdateUser(user.ID, param)
	if err != nil {
		return err
	}

	if !isUpdated {
		return echo.ErrInternalServerError
	}

	user.Name = param.Name
	user.Email = param.Email
	user.PhotoUrl = param.PhotoUrl
	sess.AddFlash("Success updating profile", "success_message")
	sess.Save(c.Request(), c.Response())
	return c.Redirect(http.StatusSeeOther, "/profile")
}

func (uc UserController) UpdateUserPassword(c echo.Context) error {
	sess, _ := session.Get("kaseer", c)

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

	if user.Email == "staff@mail.com" {
		sess.AddFlash("Cannot update this demo account", "error_message")
		sess.Save(c.Request(), c.Response())
		return c.Redirect(http.StatusSeeOther, "/profile/edit/password")
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
	sess, _ := session.Get("kaseer", c)

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
	sess, err := session.Get("kaseer", c)
	if err != nil {
		log.Println(err)
	}

	sess.Options.MaxAge = -1
	if err := sess.Save(c.Request(), c.Response()); err != nil {
		log.Println(err)
	}

	return c.Redirect(http.StatusSeeOther, "/")
}
