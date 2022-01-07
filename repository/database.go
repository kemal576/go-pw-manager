package repository

import (
	"database/sql"

	"github.com/kemal576/go-pw-manager/internal/secret"
	"github.com/kemal576/go-pw-manager/repository/login"
	"github.com/kemal576/go-pw-manager/repository/user"
)

type Database struct {
	db     *sql.DB
	users  UserRepository
	logins LoginRepository
}

func Conn() (*sql.DB, error) {
	creds, err := secret.ReadSecrets("dbsecrets")
	if err != nil {
		return nil, err
	}
	//println("postgres", "user="+creds["user"]+" password="+creds["password"]+" dbname="+creds["dbname"]+"sslmode=disable")
	db, err2 := sql.Open("postgres", " user="+creds["user"]+" password="+
		creds["password"]+" dbname="+creds["dbname"]+" sslmode=disable")
	if err2 != nil {
		return nil, err2
	}

	return db, nil
}

func New(db *sql.DB) *Database {
	return &Database{
		db:     db,
		users:  user.NewRepository(db),
		logins: login.NewRepository(db),
	}
}

// Users returns the UserRepository.
func (db *Database) Users() UserRepository {
	return db.users
}

func (db *Database) Logins() LoginRepository {
	return db.logins
}
