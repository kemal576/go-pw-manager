package api

import (
	"encoding/json"
	"net/http"

	"github.com/kemal576/go-pw-manager/models"
	"github.com/kemal576/go-pw-manager/repository"
	"golang.org/x/crypto/bcrypt"
)

func AllUsers(u repository.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := u.GetAll()
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		RespondWithJSON(w, http.StatusOK, users)
	}
}

func Create(u repository.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&user); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		pw_hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		user.Password = string(pw_hash)

		id, err := u.Create(&user)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		RespondWithJSON(w, http.StatusCreated, id)
	}
}
