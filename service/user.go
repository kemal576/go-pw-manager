package service

import (
	"github.com/kemal576/go-pw-manager/internal/app"
	"github.com/kemal576/go-pw-manager/models"
	"github.com/kemal576/go-pw-manager/utils"
)

type IUserRepository interface {
	GetAll() ([]models.User, error)
	GetById(id int) (models.User, error)
	Create(user *models.User) (int, error)
	Update(user *models.User) error
	Delete(id int) error
}

type userService struct {
	userRepository IUserRepository
}

func NewUserService(_userRepository IUserRepository) *userService {
	return &userService{userRepository: _userRepository}
}

func (u *userService) GetAll() []models.User {
	users, err := u.userRepository.GetAll()
	if err != nil {
		return nil
	}
	return users
}

func (u *userService) GetById(id int) models.User {
	user, err := u.userRepository.GetById(id)
	utils.ErrorCheck(err)
	return user
	//return models.NewDataResponse(true, "Kullanıcı Bulundu.", user)
}

func (u *userService) Create(user *models.User) int {

	hash, err := app.HashPassword(user.Password)
	utils.ErrorCheck(err)

	user.Password = hash

	id, err2 := u.userRepository.Create(user)
	utils.ErrorCheck(err2)
	return id
}

func (u *userService) Update(user *models.User) error {
	return u.userRepository.Update(user)
}

func (u *userService) Delete(id int) error {
	return u.userRepository.Delete(id)
}
