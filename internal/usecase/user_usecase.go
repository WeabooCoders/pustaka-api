package usecase

import (
	"github.com/AvinFajarF/internal/entity"
	"github.com/AvinFajarF/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase interface {
	CreateUser(username, password, alamat, email string, notelepon int) (*entity.User, error)
	Login(email, password string) (*entity.User, error)}

type userUsecase struct {
    repo repository.UserRepository
}

func NewUserUsecase(userRepository repository.UserRepository) userUsecase {
	return userUsecase{
			repo: userRepository,
		}
}

func (u *userUsecase) CreateUser(username, password, alamat, email string, notelepon int) (*entity.User, error) {
	user := &entity.User{
		Username:  username,
		Email: email,
		Password: password,
		Alamat: alamat,
		NomorTelepon: notelepon,
	}
	err := u.repo.Save(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userUsecase) Login(email, password string) (*entity.User, error) {
	user, err := u.repo.Login(email, password)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	return user, nil
}
