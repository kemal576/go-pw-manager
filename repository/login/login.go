package login

import (
	"database/sql"
	"time"

	"github.com/kemal576/go-pw-manager/models"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (u *Repository) GetAll() ([]models.Login, error) {
	response, err := u.db.Query("SELECT * FROM logins")
	if err != nil {
		return nil, err
	}

	var logins []models.Login
	for response.Next() {
		var login models.Login
		err := response.Scan(&login.Id, &login.URL, &login.Identity, &login.Password, &login.CreatedAt, &login.UpdatedAt, &login.UserId)
		if err != nil {
			return logins, err
		}
		logins = append(logins, login)
	}
	return logins, nil
}

func (u *Repository) GetById(id int) (models.Login, error) {
	var login models.Login
	response, err := u.db.Query("SELECT * FROM logins WHERE id=$1", id)
	if err != nil {
		return login, err
	}
	for response.Next() {
		err := response.Scan(&login.Id, &login.URL, &login.Identity, &login.Password, &login.CreatedAt, &login.UpdatedAt, &login.UserId)
		if err != nil {
			return login, err
		}
	}
	return login, nil
}

func (u *Repository) GetLoginsByUserId(userId int) ([]models.Login, error) {
	var logins []models.Login
	response, err := u.db.Query("SELECT * FROM logins WHERE user_id=$1", userId)
	if err != nil {
		return logins, err
	}

	for response.Next() {
		var login models.Login
		err := response.Scan(&login.Id, &login.URL, &login.Identity, &login.Password, &login.CreatedAt, &login.UpdatedAt, &login.UserId)
		if err != nil {
			return logins, err
		}
		logins = append(logins, login)
	}
	return logins, nil
}

func (u *Repository) GetLoginByUrl(userId int, url string) (models.Login, error) {
	var login models.Login
	response, err := u.db.Query("SELECT * FROM logins WHERE user_id=$1 AND url=$2", userId, url)
	if err != nil {
		return login, err
	}

	err = response.Scan(&login.Id, &login.URL, &login.Identity, &login.Password, &login.CreatedAt, &login.UpdatedAt, &login.UserId)
	if err != nil {
		return login, err
	}
	return login, nil
}

func (u *Repository) Create(login *models.Login) (int, error) {
	var lastInsertId int
	stmt, err := u.db.Prepare("INSERT INTO logins(url,identity,password,user_id) VALUES($1,$2,$3,$4) returning id;")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	stmt.QueryRow(login.URL, login.Identity, login.Password, login.UserId).Scan(&lastInsertId)
	return lastInsertId, nil
}

func (u *Repository) Update(login *models.Login) error {
	stmt, err := u.db.Prepare("UPDATE logins SET url=$1, identity=$2, password=$3, updated_at=$4 WHERE id=$5")
	if err != nil {
		return err
	}
	defer stmt.Close()
	stmt.Exec(login.URL, login.Identity, login.Password, time.Now(), login.Id)

	return nil
}

func (u *Repository) Delete(id int) error {
	stmt, err := u.db.Prepare("DELETE FROM logins WHERE id=$1")
	if err != nil {
		return err
	}
	defer stmt.Close()
	stmt.Exec(id)
	return nil
}
