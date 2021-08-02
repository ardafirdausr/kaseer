package mysql

import (
	"context"
	"database/sql"
	"log"

	"github.com/ardafirdausr/kaseer/internal/entity"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(DB *sql.DB) *UserRepository {
	return &UserRepository{DB: DB}
}

func (repo UserRepository) GetUserByID(ctx context.Context, ID int64) (*entity.User, error) {
	var row *sql.Row
	query := "SELECT * FROM users WHERE id = ?"
	if tx, ok := ctx.Value(MySQLTransactionKey("tx")).(*sql.Tx); ok {
		row = tx.QueryRow(query, ID)
	} else {
		row = repo.DB.QueryRowContext(ctx, query, ID)
	}

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

func (repo UserRepository) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	var row *sql.Row
	query := "SELECT * FROM users WHERE email = ?"
	if tx, ok := ctx.Value(MySQLTransactionKey("tx")).(*sql.Tx); ok {
		row = tx.QueryRow(query, email)
	} else {
		row = repo.DB.QueryRowContext(ctx, query, email)
	}

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

func (repo UserRepository) UpdateByID(ctx context.Context, ID int64, param entity.UpdateUserParam) (bool, error) {
	query := "UPDATE users SET name = ?, email = ?, photo_url = ? WHERE id = ?"
	var err error
	if tx, ok := ctx.Value(MySQLTransactionKey("tx")).(*sql.Tx); ok {
		_, err = tx.Exec(query, param.Name, param.Email, param.PhotoUrl, ID)
	} else {
		_, err = repo.DB.ExecContext(ctx, query, param.Name, param.Email, param.PhotoUrl, ID)
	}

	if err != nil {
		log.Println(err.Error())
		return false, err
	}

	return true, nil
}

func (repo UserRepository) UpdatePasswordByID(ctx context.Context, ID int64, password string) (bool, error) {
	query := "UPDATE users SET password = ? WHERE id = ?"
	var err error
	if tx, ok := ctx.Value(MySQLTransactionKey("tx")).(*sql.Tx); ok {
		_, err = tx.Exec(query, password, ID)
	} else {
		_, err = repo.DB.ExecContext(ctx, query, password, ID)
	}

	if err != nil {
		log.Println(err.Error())
		return false, err
	}

	return true, nil
}
