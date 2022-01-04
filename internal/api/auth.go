package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/kemal576/go-pw-manager/internal/app"
	"github.com/kemal576/go-pw-manager/models"
	"github.com/kemal576/go-pw-manager/repository"
)

func IsAuthorized(endpoint func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		token, err := jwt.Parse(cookie.Value, func(t *jwt.Token) (interface{}, error) {
			key, _ := app.GetJWTSecret()
			return key, nil
		})

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
		}

		if token.Valid {
			endpoint(w, r)
		}
	})
}

func SignIn(u repository.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var signinModel models.SignIn

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&signinModel); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		user, err := u.CheckCredentials(signinModel.Email, signinModel.Password)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		token, err := app.GenerateJWT(user.Id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		userDto := models.ToUserDTO(&user)
		cookie := &http.Cookie{Name: "token", Value: token, Expires: time.Now().Add(time.Hour * 24), HttpOnly: true, Path: "/"}
		http.SetCookie(w, cookie)
		RespondWithJSON(w, http.StatusOK, userDto)
	}
}

func Signout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie := &http.Cookie{
			Name:     "token",
			Value:    "",
			Expires:  time.Unix(0, 0),
			HttpOnly: true,
			Path:     "/",
		}
		http.SetCookie(w, cookie)
		w.WriteHeader(http.StatusOK)
	}
}

func RefreshToken() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			w.WriteHeader(http.StatusBadRequest)
			return
		}

		tokenStr := cookie.Value
		claims := &models.Claims{}
		jwtKey, err := app.GetJWTSecret()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		tkn, err := jwt.ParseWithClaims(tokenStr, claims,
			func(t *jwt.Token) (interface{}, error) { return jwtKey, nil })

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if !tkn.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		expirationTime := time.Now().Add(time.Hour * 24)
		claims.ExpiresAt = expirationTime.Unix()

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(jwtKey)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		http.SetCookie(w,
			&http.Cookie{
				Name:    "refresh_token",
				Value:   tokenString,
				Expires: expirationTime,
			})
	}
}
