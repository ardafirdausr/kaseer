package entity

import "time"

type User struct {
	ID        int64
	Name      string `form:"name"`
	Email     string `form:"email"`
	PhotoUrl  *string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserCredential struct {
	Email    string `form:"email" validate:"required,email"`
	Password string `form:"password" validate:"required"`
}

type UpdateUserParam struct {
	Name     string  `db:"name" validate:"required"`
	Email    string  `db:"email" validate:"required,email"`
	PhotoUrl *string `db:"omitempty,photo_url"`
}

type UpdateUserPasswordParam struct {
	Password             string `db:"password" form:"password" validate:"required"`
	PasswordConfirmation string `form:"password_confirmation" validate:"required,eqfield=Password"`
}
