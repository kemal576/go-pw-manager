package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/kemal576/go-pw-manager/internal/app"
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

func GetUser(u repository.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := httprouter.ParamsFromContext(r.Context())
		idStr := params.ByName("id")
		if idStr == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		user, err := u.GetById(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		RespondWithJSON(w, http.StatusOK, user)
	}
}

func CreateUser(u repository.UserRepository) http.HandlerFunc {
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

		userNew, err := u.GetById(id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		RespondWithJSON(w, http.StatusCreated, userNew)
	}
}

func UpdateUser(u repository.UserRepository) http.HandlerFunc {
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

		err = u.Update(&user)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func DeleteUser(repo repository.LoginRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := httprouter.ParamsFromContext(r.Context())
		idStr := params.ByName("id")
		if idStr == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err2 := app.CheckUser(id, r)
		if err2 != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		err = repo.Delete(id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
