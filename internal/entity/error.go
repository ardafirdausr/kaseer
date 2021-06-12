package entity

type ErrNotFound struct {
	Message string
	Err     error
}

func (ent ErrNotFound) Error() string {
	return ent.Err.Error()
}

type ErrItemAlreadyExists struct {
	Message string
	Err     error
}

func (eie ErrItemAlreadyExists) Error() string {
	return eie.Err.Error()
}

type ErrValidation struct {
	Message string
	Errors  []map[string]string
	Err     error
}

func NewErrValidation(message string, errors []map[string]string, err error) *ErrValidation {
	return &ErrValidation{
		Message: message,
		Errors:  errors,
		Err:     err,
	}
}

func (ve ErrValidation) Error() string {
	if ve.Err != nil {
		return ve.Err.Error()
	}

	return ve.Message
}
