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

//This method reads database connection information from Vault and returns the newly opened connection.
func Conn() (*sql.DB, error) {
	creds, err := secret.ReadSecrets("DB_SECRETS")
	if err != nil {
		return nil, err
	}
	db, err := sql.Open("postgres", " user="+creds["USERNAME"]+" password="+
		creds["PASSWORD"]+" dbname="+creds["DB_NAME"]+" sslmode=disable")
	if err != nil {
		return nil, err
	}

	return db, nil
}

//This constructor returns the current database connection and any repositories enabled with it.
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

// Logins returns the LoginRepository.
func (db *Database) Logins() LoginRepository {
	return db.logins
}
