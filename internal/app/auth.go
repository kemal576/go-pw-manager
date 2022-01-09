package app

import (
	"net/http"
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

func CheckUser(userId int, r *http.Request) bool {
	cookieToken := r.Header["Authorization"]

	tokenStr := cookieToken[0]
	claims := &models.Claims{}
	jwtKey, err := GetJWTSecret()
	if err != nil {
		return false
	}
	tkn, err := jwt.ParseWithClaims(tokenStr, claims,
		func(t *jwt.Token) (interface{}, error) { return jwtKey, nil })

	if err != nil {
		return false
	}

	if !tkn.Valid {
		return false
	}

	return claims.UserId == userId
}

func GetJWTSecret() ([]byte, error) {
	var key []byte
	secret, err := secret.ReadSecrets("JWT")
	if err != nil {
		return key, err
	}

	key = []byte(secret["JWT_KEY"])
	return key, nil
}
