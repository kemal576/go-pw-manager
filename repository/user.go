package repository

import (
	"database/sql"
	"errors"

	"github.com/kemal576/go-pw-manager/models"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) (*userRepository, error) {
	if db == nil {
		return nil, errors.New("provided db handle to user repository is nil")
	}
	return &userRepository{db: db}, nil
}

func (u *userRepository) GetAll() ([]models.User, error) {
	response, err := u.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}

	var users []models.User
	for response.Next() {
		var user models.User
		err := response.Scan(&user.Id, &user.FirstName, &user.LastName, &user.PasswordHash, &user.CreatedAt)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (u *userRepository) GetById(id int) (models.User, error) {
	var user models.User
	response, err := u.db.Query("SELECT * FROM users WHERE id=$1", id)
	if err != nil {
		return user, err
	}
	for response.Next() {
		err := response.Scan(&user.Id, &user.FirstName, &user.LastName, &user.PasswordHash, &user.CreatedAt)
		if err != nil {
			return user, err
		}
	}
	return user, nil
}

func (u *userRepository) Create(user *models.User) (int, error) {
	var lastInsertId int
	stmt, err := u.db.Prepare("INSERT INTO users(firstname,lastname,password_hash) VALUES($1,$2,$3) returning id;")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	stmt.QueryRow(user.FirstName, user.LastName, user.PasswordHash).Scan(&lastInsertId)
	return lastInsertId, nil
}

func (u *userRepository) Update(user *models.User) error {
	stmt, err := u.db.Prepare("UPDATE users SET firstname=$1, lastname=$2, password_hash=$3 WHERE id=$4")
	if err != nil {
		return err
	}
	defer stmt.Close()
	stmt.Exec(user.FirstName, user.LastName, user.PasswordHash, user.Id)
	return nil
}

func (u *userRepository) Delete(id int) error {
	stmt, err := u.db.Prepare("DELETE FROM users WHERE id=$1")
	if err != nil {
		return err
	}
	defer stmt.Close()
	stmt.Exec(id)
	return nil
}
