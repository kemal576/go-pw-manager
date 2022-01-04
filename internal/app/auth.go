package app

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/kemal576/go-pw-manager/internal/secret"
	"github.com/kemal576/go-pw-manager/models"
)

func GenerateJWT(userId int) (string, error) {
	key, err := GetJWTSecret()
	if err != nil {
		return "", err
	}

	expirationTime := time.Now().Add(time.Hour * 24)
	claims := &models.Claims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err2 := token.SignedString(key)

	if err2 != nil {
		return "", err
	}

	return tokenString, nil
}

func GetJWTSecret() ([]byte, error) {
	var key []byte
	secret, err := secret.ReadSecrets("jwt")
	if err != nil {
		return key, err
	}

	key = []byte(secret["jwt_key"])
	return key, nil
}
