package models

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	UserId int `json:"UserId"`
	jwt.StandardClaims
}
