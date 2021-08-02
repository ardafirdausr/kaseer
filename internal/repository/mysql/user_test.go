package mysql

import (
	"context"
	"database/sql"
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ardafirdausr/kaseer/internal/entity"
	"github.com/stretchr/testify/assert"
)

func Test_GetUserByID_Failed_WhenNoRows(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	ctx := context.TODO()
	userID := int64(1)
	query := regexp.QuoteMeta("SELECT * FROM users WHERE id = ?")
	mock.ExpectQuery(query).
		WithArgs(userID).
		WillReturnError(sql.ErrNoRows)

	userRepository := NewUserRepository(db)
	user, err := userRepository.GetUserByID(ctx, userID)
	assert.NotNil(t, err)
	assert.IsType(t, entity.ErrNotFound{}, err)
	assert.Nil(t, user)
}

func Test_GetUserByID_Failed_WhenErrorOccurs(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	ctx := context.TODO()
	userID := int64(1)
	query := regexp.QuoteMeta("SELECT * FROM users WHERE id = ?")

	mock.ExpectQuery(query).
		WithArgs(userID).
		WillReturnError(errors.New("failed get users"))

	userRepository := NewUserRepository(db)
	user, err := userRepository.GetUserByID(ctx, userID)
	assert.NotNil(t, err)
	assert.Nil(t, user)
}

func Test_GetUserByID_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	ctx := context.TODO()
	eUser := entity.User{
		ID:        1,
		Name:      "John Doe",
		Email:     "JohnDoe@mail.com",
		PhotoUrl:  nil,
		Password:  "secret",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	query := regexp.QuoteMeta("SELECT * FROM users WHERE id = ?")
	rows := sqlmock.NewRows([]string{"id", "name", "email", "photo_url", "passsword", "created_at", "updated_at"})
	rows.AddRow(eUser.ID, eUser.Name, eUser.Email, eUser.PhotoUrl, eUser.Password, eUser.CreatedAt, eUser.UpdatedAt)
	mock.ExpectQuery(query).
		WithArgs(eUser.ID).
		WillReturnRows(rows)
	userRepository := NewUserRepository(db)
	aUser, err := userRepository.GetUserByID(ctx, eUser.ID)
	assert.Nil(t, err)
	assert.ObjectsAreEqualValues(eUser, aUser)
}

func Test_GetUserByEmail_Failed_WhenNoRows(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	ctx := context.TODO()
	userEmail := "some@mail.com"
	query := regexp.QuoteMeta("SELECT * FROM users WHERE email = ?")
	mock.ExpectQuery(query).
		WithArgs(userEmail).
		WillReturnError(sql.ErrNoRows)

	userRepository := NewUserRepository(db)
	user, err := userRepository.GetUserByEmail(ctx, userEmail)
	assert.NotNil(t, err)
	assert.IsType(t, entity.ErrNotFound{}, err)
	assert.Nil(t, user)
}

func Test_GetUserByEmail_Failed_WhenErrorOccurs(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	ctx := context.TODO()
	userEmail := "some@mail.com"
	query := regexp.QuoteMeta("SELECT * FROM users WHERE email = ?")

	mock.ExpectQuery(query).
		WithArgs(userEmail).
		WillReturnError(errors.New("failed get users"))

	userRepository := NewUserRepository(db)
	user, err := userRepository.GetUserByEmail(ctx, userEmail)
	assert.NotNil(t, err)
	assert.Nil(t, user)
}

func Test_GetUserByEmail_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	ctx := context.TODO()
	eUser := entity.User{
		ID:        1,
		Name:      "John Doe",
		Email:     "JohnDoe@mail.com",
		PhotoUrl:  nil,
		Password:  "secret",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	query := regexp.QuoteMeta("SELECT * FROM users WHERE email = ?")
	rows := sqlmock.NewRows([]string{"id", "name", "email", "photo_url", "passsword", "created_at", "updated_at"})
	rows.AddRow(eUser.ID, eUser.Name, eUser.Email, eUser.PhotoUrl, eUser.Password, eUser.CreatedAt, eUser.UpdatedAt)
	mock.ExpectQuery(query).
		WithArgs(eUser.Email).
		WillReturnRows(rows)
	userRepository := NewUserRepository(db)
	aUser, err := userRepository.GetUserByEmail(ctx, eUser.Email)
	assert.Nil(t, err)
	assert.ObjectsAreEqualValues(eUser, aUser)
}

func Test_UpdateByID_Failed(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	ctx := context.TODO()
	userID := int64(1)
	param := entity.UpdateUserParam{
		Name:     "John",
		Email:    "JohnDoe@mail.com",
		PhotoUrl: nil,
	}
	query := regexp.QuoteMeta("UPDATE users SET name = ?, email = ?, photo_url = ? WHERE id = ?")
	mock.ExpectExec(query).
		WithArgs(param.Name, param.Email, param.PhotoUrl, userID).
		WillReturnError(errors.New("failed to update user"))

	userRepository := NewUserRepository(db)
	updated, err := userRepository.UpdateByID(ctx, userID, param)
	assert.NotNil(t, err)
	assert.False(t, updated)
}

func Test_UpdateByID_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	ctx := context.TODO()
	userID := int64(1)
	param := entity.UpdateUserParam{
		Name:     "John",
		Email:    "JohnDoe@mail.com",
		PhotoUrl: nil,
	}
	query := regexp.QuoteMeta("UPDATE users SET name = ?, email = ?, photo_url = ? WHERE id = ?")
	mock.ExpectExec(query).
		WithArgs(param.Name, param.Email, param.PhotoUrl, userID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	userRepository := NewUserRepository(db)
	updated, err := userRepository.UpdateByID(ctx, userID, param)
	assert.Nil(t, err)
	assert.True(t, updated)
}

func Test_UpdatePasswordByID_Failed(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	ctx := context.TODO()
	userID := int64(1)
	password := "new-secret-pass"
	query := regexp.QuoteMeta("UPDATE users SET password = ? WHERE id = ?")
	mock.ExpectExec(query).
		WithArgs(password, userID).
		WillReturnError(errors.New("failed to update user"))

	userRepository := NewUserRepository(db)
	updated, err := userRepository.UpdatePasswordByID(ctx, userID, password)
	assert.NotNil(t, err)
	assert.False(t, updated)
}

func Test_UpdatePasswordByID_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	ctx := context.TODO()
	userID := int64(1)
	password := "new-secret-pass"
	query := regexp.QuoteMeta("UPDATE users SET password = ? WHERE id = ?")
	mock.ExpectExec(query).
		WithArgs(password, userID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	userRepository := NewUserRepository(db)
	updated, err := userRepository.UpdatePasswordByID(ctx, userID, password)
	assert.Nil(t, err)
	assert.True(t, updated)
}
