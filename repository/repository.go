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
