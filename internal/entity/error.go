package entity

import "net/http"

type ErrNotFound struct {
	Message string
	Err     error
}

func (ent ErrNotFound) Error() string {
	if ent.Err != nil {
		return ent.Err.Error()
	} else if ent.Message != "" {
		return ent.Message
	} else {
		return http.StatusText(http.StatusBadGateway)
	}
}

type ErrItemAlreadyExists struct {
	Message string
	Err     error
}

func (eie ErrItemAlreadyExists) Error() string {
	if eie.Err != nil {
		return eie.Err.Error()
	} else if eie.Message != "" {
		return eie.Message
	} else {
		return http.StatusText(http.StatusBadRequest)
	}
}

type ErrValidation struct {
	Message string
	Errors  map[string]string
}

func (ve ErrValidation) Error() string {
	if ve.Message != "" {
		return ve.Message
	} else {
		return http.StatusText(http.StatusBadRequest)
	}
}
