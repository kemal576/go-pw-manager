package models

import "time"

type User struct {
	Id        int       `json:"id"`
	FirstName string    `json:"firstname"`
	LastName  string    `json:"lastname"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	Email     string    `json:"email"`
}

type UserDTO struct {
	Id        int    `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
}

type UserReturnDTO struct {
	Id        int       `json:"id"`
	FirstName string    `json:"firstname"`
	LastName  string    `json:"lastname"`
	CreatedAt time.Time `json:"created_at"`
	Email     string    `json:"email"`
}

func ToUserDTO(user *User) *UserDTO {
	return &UserDTO{
		Id:        user.Id,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}
}

func ToUserReturnDTO(user *User) *UserReturnDTO {
	return &UserReturnDTO{
		Id:        user.Id,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}
}
