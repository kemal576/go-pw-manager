package api

import (
	"encoding/json"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/kemal576/go-pw-manager/internal/app"
	"github.com/kemal576/go-pw-manager/models"
	"github.com/kemal576/go-pw-manager/repository"
)

//This method checks the request header of the sent endpoint method and if there is a jwt token,
//it returns results by checking the validity of the token.
//If the token is valid, the current endpoint is executed. If not valid, the specified response is returned.
func IsAuthorized(endpoint func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Authorization"] == nil {
			RespondWithError(w, http.StatusUnauthorized, "Authorization header is empty!")
			return
		}
		cookieToken := r.Header["Authorization"]
		token, err := jwt.Parse(cookieToken[0], func(t *jwt.Token) (interface{}, error) {
			key, _ := app.GetJWTSecret()
			return key, nil
		})

		if err != nil {
			RespondWithError(w, http.StatusUnauthorized, "JWT is not valid!")
		}

		if token.Valid {
			endpoint(w, r)
		}
	})
}

//This method compares the submitted credentials with those in the database.
//Returns the JWT token if the information is valid.
//If the information is not valid, it returns the specified response.
func SignIn(u repository.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var signinModel models.SignIn
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&signinModel); err != nil {
			RespondWithError(w, http.StatusUnprocessableEntity, "Payload could not be parsed!")
			return
		}
		defer r.Body.Close()

		user, err := u.CheckCredentials(signinModel.Email, signinModel.Password)
		if err != nil {
			RespondWithError(w, http.StatusUnauthorized, "Email and password does not match!")
			return
		}

		token, err := app.GenerateJWT(user.Id)
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, "Error occurred while generating jwt.!")
			return
		}

		tokenRes := &models.TokenResponse{UserId: user.Id, Email: user.Email, TokenString: token}
		RespondWithJSON(w, http.StatusOK, tokenRes)
	}
}

/*
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
*/

/*
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
}*/
