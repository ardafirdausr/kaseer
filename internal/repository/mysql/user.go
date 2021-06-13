package mysql

import (
	"database/sql"
	"log"

	"github.com/ardafirdausr/go-pos/internal/entity"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(DB *sql.DB) *UserRepository {
	return &UserRepository{DB: DB}
}

func (ur UserRepository) GetUserByID(ID int64) (*entity.User, error) {
	row := ur.DB.QueryRow("SELECT * FROM users WHERE id = ?", ID)

	var user entity.User
	var err = row.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.PhotoUrl,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		err := entity.ErrNotFound{
			Message: "User not found",
			Err:     err,
		}
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur UserRepository) GetUserByEmail(email string) (*entity.User, error) {
	row := ur.DB.QueryRow("SELECT * FROM users WHERE email = ?", email)
	var user entity.User
	var err = row.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.PhotoUrl,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		err := entity.ErrNotFound{
			Message: "User not found",
			Err:     err,
		}
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur UserRepository) UpdateByID(ID int64, param entity.UpdateUserParam) (bool, error) {
	_, err := ur.DB.Exec(
		"UPDATE users SET name = ?, email = ?, photo_url = ? WHERE id = ?",
		param.Name, param.Email, param.PhotoUrl, ID)
	if err != nil {
		log.Println(err.Error())
		return false, err
	}

	return true, nil
}

func (ur UserRepository) UpdatePasswordByID(ID int64, password string) (bool, error) {
	_, err := ur.DB.Exec(
		"UPDATE users SET password = ? WHERE id = ?", password, ID)
	if err != nil {
		log.Println(err.Error())
		return false, err
	}

	return true, nil
}
