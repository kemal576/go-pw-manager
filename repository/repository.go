package repository

import "github.com/kemal576/go-pw-manager/models"

type UserRepository interface {
	GetAll() ([]models.User, error)
	GetById(id int) (models.User, error)
	GetByEmail(email string) (models.User, error)
	CheckCredentials(email, password string) (models.User, error)
	Create(user *models.User) (int, error)
	Update(user *models.User) error
	Delete(id int) error
}

type LoginRepository interface {
	GetAll() ([]models.Login, error)
	GetById(id int) (models.Login, error)
	GetLoginsByUserId(userId int) ([]models.Login, error)
	GetLoginByUrl(userId int, url string) (models.Login, error)
	Create(login *models.Login) (int, error)
	Update(login *models.Login) error
	Delete(id int) error
}
