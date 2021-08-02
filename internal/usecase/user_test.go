package usecase

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/ardafirdausr/kaseer/internal/entity"
	"github.com/ardafirdausr/kaseer/internal/mocks"
	"github.com/stretchr/testify/assert"
)

var user = entity.User{
	ID:        1,
	Name:      "John Doe",
	Email:     "johndoe@mail.com",
	PhotoUrl:  nil,
	Password:  "soMeRandomPwd",
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
}

func Test_GetUserByID_Failed(t *testing.T) {
	ctx := context.TODO()

	mockUserRepository := new(mocks.UserRepository)
	mockUserRepository.On("GetUserByID", ctx, user.ID).Return(nil, errors.New("failed get user by id"))
	mockStorage := new(mocks.Storage)

	userUsecase := NewUserUsecase(mockUserRepository, mockStorage)
	user, err := userUsecase.GetUserByID(ctx, user.ID)
	assert.NotNil(t, err)
	assert.Nil(t, user)
}

func Test_GetUserByID_Success(t *testing.T) {
	ctx := context.TODO()

	mockUserRepository := new(mocks.UserRepository)
	mockUserRepository.On("GetUserByID", ctx, user.ID).Return(&user, nil)
	mockStorage := new(mocks.Storage)

	userUsecase := NewUserUsecase(mockUserRepository, mockStorage)
	actualUser, err := userUsecase.GetUserByID(ctx, user.ID)
	assert.Nil(t, err)
	assert.ObjectsAreEqualValues(actualUser, user)
}

func Test_GetUserByCredential_Failed_WhenGettingUser(t *testing.T) {
	ctx := context.TODO()
	credential := entity.UserCredential{
		Email:    user.Email,
		Password: user.Password,
	}

	mockUserRepository := new(mocks.UserRepository)
	mockUserRepository.On("GetUserByEmail", ctx, credential.Email).Return(nil, errors.New("failed get user by id"))
	mockStorage := new(mocks.Storage)

	userUsecase := NewUserUsecase(mockUserRepository, mockStorage)
	user, err := userUsecase.GetUserByCredential(ctx, credential)
	assert.NotNil(t, err)
	assert.Nil(t, user)
}

func Test_GetUserByCredential_Failed_WhenPasswordNotEqual(t *testing.T) {
	oriHash := stringsHash
	stringsHash = func(v string) string {
		return "differentPassword"
	}

	ctx := context.TODO()
	credential := entity.UserCredential{
		Email:    user.Email,
		Password: user.Password,
	}
	mockUserRepository := new(mocks.UserRepository)
	mockUserRepository.On("GetUserByEmail", ctx, credential.Email).Return(&user, nil)
	mockStorage := new(mocks.Storage)

	userUsecase := NewUserUsecase(mockUserRepository, mockStorage)
	aUser, err := userUsecase.GetUserByCredential(ctx, credential)
	assert.NotNil(t, err)
	assert.Nil(t, aUser)
	stringsHash = oriHash
}

func Test_GetUserByCredential_Success(t *testing.T) {
	oriHash := stringsHash
	stringsHash = func(v string) string {
		return user.Password
	}

	ctx := context.TODO()
	credential := entity.UserCredential{
		Email:    user.Email,
		Password: user.Password,
	}
	mockUserRepository := new(mocks.UserRepository)
	mockUserRepository.On("GetUserByEmail", ctx, credential.Email).Return(&user, nil)
	mockStorage := new(mocks.Storage)

	userUsecase := NewUserUsecase(mockUserRepository, mockStorage)
	aUser, err := userUsecase.GetUserByCredential(ctx, credential)
	assert.Nil(t, err)
	assert.ObjectsAreEqualValues(user, aUser)
	stringsHash = oriHash
}

func Test_UpdateUser_Failed(t *testing.T) {
	ctx := context.TODO()
	updateParam := entity.UpdateUserParam{
		Name:     "John doie",
		Email:    "newMail@mail.com",
		PhotoUrl: nil,
	}

	mockUserRepository := new(mocks.UserRepository)
	mockUserRepository.On("UpdateByID", ctx, user.ID, updateParam).Return(false, errors.New("failed to update the user"))
	mockStorage := new(mocks.Storage)

	userUsecase := NewUserUsecase(mockUserRepository, mockStorage)
	isUpdated, err := userUsecase.UpdateUser(ctx, user.ID, updateParam)
	assert.NotNil(t, err)
	assert.False(t, isUpdated)
}

func Test_UpdateUser_Success(t *testing.T) {
	ctx := context.TODO()
	updateParam := entity.UpdateUserParam{
		Name:     "John doie",
		Email:    "newMail@mail.com",
		PhotoUrl: nil,
	}

	mockUserRepository := new(mocks.UserRepository)
	mockUserRepository.On("UpdateByID", ctx, user.ID, updateParam).Return(true, nil)
	mockStorage := new(mocks.Storage)

	userUsecase := NewUserUsecase(mockUserRepository, mockStorage)
	isUpdated, err := userUsecase.UpdateUser(ctx, user.ID, updateParam)
	assert.Nil(t, err)
	assert.True(t, isUpdated)
}

func Test_UpdateUserPassword_Failed(t *testing.T) {
	oriHash := stringsHash
	stringsHash = func(v string) string {
		return "hashedPassword"
	}

	ctx := context.TODO()
	mockUserRepository := new(mocks.UserRepository)
	mockUserRepository.On("UpdatePasswordByID", ctx, user.ID, "hashedPassword").Return(false, errors.New("failed to update the user"))
	mockStorage := new(mocks.Storage)

	userUsecase := NewUserUsecase(mockUserRepository, mockStorage)
	isUpdated, err := userUsecase.UpdateUserPassword(ctx, user.ID, "new-password")
	assert.NotNil(t, err)
	assert.False(t, isUpdated)
	stringsHash = oriHash
}

func Test_UpdateUserPassword_Success(t *testing.T) {
	oriHash := stringsHash
	stringsHash = func(v string) string {
		return "hashedPassword"
	}

	ctx := context.TODO()
	mockUserRepository := new(mocks.UserRepository)
	mockUserRepository.On("UpdatePasswordByID", ctx, user.ID, "hashedPassword").Return(true, nil)
	mockStorage := new(mocks.Storage)

	userUsecase := NewUserUsecase(mockUserRepository, mockStorage)
	isUpdated, err := userUsecase.UpdateUserPassword(ctx, user.ID, "new-password")
	assert.Nil(t, err)
	assert.True(t, isUpdated)
	stringsHash = oriHash
}
