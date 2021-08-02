package usecase

import (
	"context"
	"fmt"
	"log"
	"mime/multipart"
	"path/filepath"

	"github.com/ardafirdausr/kaseer/internal"
	"github.com/ardafirdausr/kaseer/internal/entity"
	"github.com/ardafirdausr/kaseer/internal/pkg/strings"
)

var stringsHash = strings.Hash

type UserUsecase struct {
	userRepository internal.UserRepository
	storage        internal.Storage
}

func NewUserUsecase(userRepository internal.UserRepository, storage internal.Storage) *UserUsecase {
	return &UserUsecase{userRepository, storage}
}

func (uu UserUsecase) GetUserByID(ctx context.Context, ID int64) (*entity.User, error) {
	user, err := uu.userRepository.GetUserByID(ctx, ID)
	if err != nil {
		log.Println(err.Error())
	}

	return user, err
}

func (uu UserUsecase) GetUserByCredential(ctx context.Context, credential entity.UserCredential) (*entity.User, error) {
	user, err := uu.userRepository.GetUserByEmail(ctx, credential.Email)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	hashedPassword := stringsHash(credential.Password)
	isPasswordEqual := hashedPassword == user.Password
	if !isPasswordEqual {
		err := entity.ErrInvalidCredential{
			Message: "Invalid Password",
			Err:     nil,
		}
		return nil, err
	}

	return user, nil
}

func (uu UserUsecase) SaveUserPhoto(ctx context.Context, user *entity.User, photo *multipart.FileHeader) (string, error) {
	photoName := fmt.Sprintf("user-%d", user.ID)
	photoExt := filepath.Ext(photo.Filename)
	filename := photoName + photoExt
	photoDirectory := filepath.Join("image", "user")
	return uu.storage.Save(photo, photoDirectory, filename)
}

func (uu UserUsecase) UpdateUser(ctx context.Context, ID int64, param entity.UpdateUserParam) (bool, error) {
	isUpdated, err := uu.userRepository.UpdateByID(ctx, ID, param)
	if err != nil {
		log.Println(err.Error())
		return false, err
	}

	return isUpdated, nil
}

func (uu UserUsecase) UpdateUserPassword(ctx context.Context, ID int64, password string) (bool, error) {
	hashedPassword := stringsHash(password)
	isUpdated, err := uu.userRepository.UpdatePasswordByID(ctx, ID, hashedPassword)
	if err != nil {
		log.Println(err.Error())
	}

	return isUpdated, err
}
