package models

type TokenResponse struct {
	UserId      int    `json:"user_id"`
	Email       string `json:"email"`
	TokenString string `json:"token"`
}
