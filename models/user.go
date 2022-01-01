package models

import "time"

type User struct {
	Id        int
	FirstName string
	LastName  string
	Password  string
	CreatedAt time.Time
	Email     string
}
