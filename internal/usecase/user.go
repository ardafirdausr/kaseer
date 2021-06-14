package usecase

import (
	"crypto/sha1"
	"fmt"
	"log"
	"mime/multipart"
	"path/filepath"

	"github.com/ardafirdausr/go-pos/internal"
	"github.com/ardafirdausr/go-pos/internal/entity"
)

type UserUsecase struct {
	userRepository internal.UserRepository
}

func NewUserUsecase(userRepository internal.UserRepository) *UserUsecase {
	return &UserUsecase{userRepository: userRepository}
}

func (uu UserUsecase) GetUserByID(ID int64) (*entity.User, error) {
	user, err := uu.userRepository.GetUserByID(ID)
	if err != nil {
		log.Println(err.Error())
	}

	return user, err
}

func (uu UserUsecase) GetUserByCredential(credential entity.UserCredential) (*entity.User, error) {
	user, err := uu.userRepository.GetUserByEmail(credential.Email)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	hash := sha1.New()
	hash.Write([]byte(credential.Password))
	hashed := hash.Sum(nil)
	isPasswordEqual := fmt.Sprintf("%x", hashed) == user.Password
	if !isPasswordEqual {
		err := entity.ErrInvalidCredential{
			Message: "Invalid Password",
			Err:     nil,
		}
		return nil, err
	}

	return user, nil
}

func (uu UserUsecase) SaveUserPhoto(storage internal.Storage, user *entity.User, photo *multipart.FileHeader) (string, error) {
	photoName := fmt.Sprintf("user-%d", user.ID)
	photoExt := filepath.Ext(photo.Filename)
	filename := photoName + photoExt
	photoDirectory := filepath.Join("image", "user")
	return storage.Save(photo, photoDirectory, filename)
}

func (uu UserUsecase) UpdateUser(ID int64, param entity.UpdateUserParam) (bool, error) {
	isUpdated, err := uu.userRepository.UpdateByID(ID, param)
	if err != nil {
		log.Println(err.Error())
		return false, err
	}

	return isUpdated, nil
}

func (uu UserUsecase) UpdateUserPassword(ID int64, password string) (bool, error) {
	hash := sha1.New()
	hash.Write([]byte(password))
	hashBytePass := hash.Sum(nil)
	hashedPassword := fmt.Sprintf("%x", hashBytePass)

	isUpdated, err := uu.userRepository.UpdatePasswordByID(ID, hashedPassword)
	if err != nil {
		log.Println(err.Error())
	}

	return isUpdated, err
}
