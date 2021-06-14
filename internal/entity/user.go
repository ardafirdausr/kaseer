package entity

import "time"

type User struct {
	ID        int64
	Name      string    `db:"name"`
	Email     string    `db:"email"`
	PhotoUrl  *string   `db:"omitempty,photo_url"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type UserCredential struct {
	Email    string `form:"email" validate:"required,email"`
	Password string `form:"password" validate:"required"`
}

type UpdateUserParam struct {
	Name     string  `db:"name"`
	Email    string  `db:"email" validate:"email"`
	PhotoUrl *string `db:"omitempty,photo_url" validate:"url"`
}

type UpdateUserPasswordParam struct {
	Password             string `db:"password" form:"password" validate:"required"`
	PasswordConfirmation string `form:"password_confirmation" validate:"required,eqfield=Password"`
}
