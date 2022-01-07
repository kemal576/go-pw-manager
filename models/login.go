package models

import "time"

type Login struct {
	Id        int       `json:"id"`
	URL       string    `json:"url"`
	Identity  string    `json:"identity"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserId    int       `json:"userId"`
}

type LoginDTO struct {
	URL      string `json:"url"`
	Identity string `json:"identity"`
	Password string `json:"password"`
	UserId   int    `json:"userId"`
}

func ToLogin(loginDto LoginDTO) *Login {
	return &Login{
		URL:      loginDto.URL,
		Identity: loginDto.Identity,
		Password: loginDto.Password,
		UserId:   loginDto.UserId,
	}
}
