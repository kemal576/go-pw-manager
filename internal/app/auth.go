package app

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/kemal576/go-pw-manager/internal/secret"
	"github.com/kemal576/go-pw-manager/models"
)

//This method generates and returns a new JWT by receiving the sent user information.
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

//This method compares the sent user id with the user registered to the token in the request.
//In this way, it prevents unauthorized access.
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

//This method reads and returns the jwt secret key stored in the Vault.
func GetJWTSecret() ([]byte, error) {
	var key []byte
	secret, err := secret.ReadSecrets("JWT")
	if err != nil {
		return key, err
	}

	key = []byte(secret["JWT_KEY"])
	return key, nil
}
