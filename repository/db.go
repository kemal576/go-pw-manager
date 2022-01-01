package repository

import (
	"database/sql"

	"github.com/kemal576/go-pw-manager/internal/secret"
)

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

/*
func CheckVersion() string {
	db, err := sql.Open("mysql", "root:1234567@tcp(127.0.0.1:3306)/crud")
	utils.ErrorCheck(err)

	var version string
	err2 := db.QueryRow("SELECT VERSION()").Scan(&version)
	utils.ErrorCheck(err2)
	return version
}
*/
