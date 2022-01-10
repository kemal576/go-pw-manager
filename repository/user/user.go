package user

import (
	"database/sql"

	"github.com/kemal576/go-pw-manager/models"
	"golang.org/x/crypto/bcrypt"
)

type Repository struct {
	db *sql.DB
}

//This method returns a new UserRepository using the connection it received.
func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

//This method returns all users found in the database.
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

//This method returns the user registered in the database with the id sent with the parameter.
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

//This method returns the user registered in the database with the email sent with the parameter.
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

//This method compares the submitted user credentials with the information in the database.
//Returns user data if the information matches. If it doesn't match, it returns an error.
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

//This method saves the sent user data to the database and returns the new user's id information.
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

//This method updates the information of the user registered in the database with the information sent.
func (u *Repository) Update(user *models.User) error {
	stmt, err := u.db.Prepare("UPDATE users SET firstname=$1, lastname=$2, password_hash=$3 WHERE id=$4")
	if err != nil {
		return err
	}
	defer stmt.Close()
	stmt.Exec(user.FirstName, user.LastName, user.Password, user.Id)
	return nil
}

//This method deletes the registered user from the database with the id information sent.
func (u *Repository) Delete(id int) error {
	stmt, err := u.db.Prepare("DELETE FROM users WHERE id=$1")
	if err != nil {
		return err
	}
	defer stmt.Close()
	stmt.Exec(id)
	return nil
}
