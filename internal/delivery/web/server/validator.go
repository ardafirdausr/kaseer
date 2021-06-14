package server

import (
	"fmt"
	"net/http"

	"github.com/ardafirdausr/go-pos/internal/entity"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (v *CustomValidator) Validate(i interface{}) error {
	// paramReflectValue := reflect.ValueOf(i)
	// switch paramReflectValue.Kind() {
	// case reflect.Map:
	// 	i, ok = i.(echo.Map)
	// 	if !ok {

	// 	}
	// 	v.validator.ValidateMap(i["data"], i["rules"])
	// case reflect.Struct:
	// 	v.validator.Struct(i)
	// }

	err := v.validator.Struct(i)
	if _, ok := err.(*validator.InvalidValidationError); ok {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	} else if validationErrors, ok := err.(validator.ValidationErrors); ok {
		verr := entity.ErrValidation{
			Message: "Invalid format data",
			Errors:  map[string]string{},
		}
		for _, validationError := range validationErrors {
			fieldName := validationError.Field()
			paramValue := validationError.Param()
			errMessage := ""
			switch validationError.Tag() {
			case "required":
				errMessage = fmt.Sprintf("%s is required", fieldName)
			case "gt":
				errMessage = fmt.Sprintf("Value of %s must be greater than %s", fieldName, paramValue)
			case "gte":
				errMessage = fmt.Sprintf("Value of %s must be greater or equal to %s", fieldName, paramValue)
			case "email":
				errMessage = fmt.Sprintf("Value of %s must be valid email", fieldName)
			case "eqfield":
				errMessage = fmt.Sprintf("Value of %s must be equal with", paramValue)
			}
			verr.Errors[fieldName] = errMessage
		}
		return verr
	}

	return nil
}
