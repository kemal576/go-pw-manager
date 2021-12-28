package repository

import (
	"database/sql"

	"github.com/kemal576/go-pw-manager/utils"
)

func Conn() *sql.DB {
	db, err := sql.Open("postgres", "user=postgres password=12345 dbname=pw-manager sslmode=disable")
	utils.ErrorCheck(err)
	return db
}

func CheckVersion() string {
	db, err := sql.Open("mysql", "root:1234567@tcp(127.0.0.1:3306)/crud")
	utils.ErrorCheck(err)

	var version string
	err2 := db.QueryRow("SELECT VERSION()").Scan(&version)
	utils.ErrorCheck(err2)
	return version
}
