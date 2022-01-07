package user

import (
	"database/sql"

	"github.com/kemal576/go-pw-manager/models"
	"golang.org/x/crypto/bcrypt"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (u *Repository) GetAll() ([]models.User, error) {
	response, err := u.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}

	var users []models.User
	for response.Next() {
		var user models.User
		err := response.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Password, &user.CreatedAt, &user.Email)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (u *Repository) GetById(id int) (models.User, error) {
	var user models.User
	response, err := u.db.Query("SELECT * FROM users WHERE id=$1", id)
	if err != nil {
		return user, err
	}
	for response.Next() {
		err := response.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Password, &user.CreatedAt, &user.Email)
		if err != nil {
			return user, err
		}
	}
	return user, nil
}

func (u *Repository) GetByEmail(email string) (models.User, error) {
	var user models.User
	response, err := u.db.Query("SELECT * FROM users WHERE email=$1", email)
	if err != nil {
		return user, err
	}
	for response.Next() {
		err := response.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Password, &user.CreatedAt, &user.Email)
		if err != nil {
			return user, err
		}
	}
	return user, nil
}

func (u *Repository) CheckCredentials(email, password string) (models.User, error) {
	user, err := u.GetByEmail(email)
	if err != nil {
		return user, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}

func (u *Repository) Create(user *models.User) (int, error) {
	var lastInsertId int
	stmt, err := u.db.Prepare("INSERT INTO users(firstname,lastname,password_hash,email) VALUES($1,$2,$3,$4) returning id;")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	stmt.QueryRow(user.FirstName, user.LastName, user.Password, user.Email).Scan(&lastInsertId)
	return lastInsertId, nil
}

func (u *Repository) Update(user *models.User) error {
	stmt, err := u.db.Prepare("UPDATE users SET firstname=$1, lastname=$2, password_hash=$3 WHERE id=$4")
	if err != nil {
		return err
	}
	defer stmt.Close()
	stmt.Exec(user.FirstName, user.LastName, user.Password, user.Id)
	return nil
}

func (u *Repository) Delete(id int) error {
	stmt, err := u.db.Prepare("DELETE FROM users WHERE id=$1")
	if err != nil {
		return err
	}
	defer stmt.Close()
	stmt.Exec(id)
	return nil
}
