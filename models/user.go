package models

import "time"

type User struct {
	Id           int
	FirstName    string
	LastName     string
	PasswordHash string
	CreatedAt    time.Time
}
