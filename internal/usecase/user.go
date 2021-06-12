package usecase

import (
	"crypto/sha1"
	"fmt"
	"log"

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

func (uu UserUsecase) GetUserByCredential(email, password string) (*entity.User, error) {
	user, err := uu.userRepository.GetUserByEmail(email)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	hash := sha1.New()
	hash.Write([]byte(password))
	hashed := hash.Sum(nil)
	isPasswordEqual := fmt.Sprintf("%x", hashed) == user.Password
	if !isPasswordEqual {
		return nil, err
	}

	return user, nil
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
	hashedPassword := string(hash.Sum(nil))

	isUpdated, err := uu.userRepository.UpdatePasswordByID(ID, hashedPassword)
	if err != nil {
		log.Println(err.Error())
	}

	return isUpdated, err
}
