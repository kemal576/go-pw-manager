package api

import (
	"encoding/json"
	"net/http"

	"github.com/kemal576/go-pw-manager/internal/app"
	"github.com/kemal576/go-pw-manager/models"
	"github.com/kemal576/go-pw-manager/repository"
	"golang.org/x/crypto/bcrypt"
)

func AllUsers(u repository.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := u.GetAll()
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, "Something went wrong about DB!")
		}

		if len(users) == 0 {
			RespondWithError(w, http.StatusNotFound, "No users found!")
		}

		RespondWithJSON(w, http.StatusOK, users)
	}
}

func GetUser(u repository.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		/*params := httprouter.ParamsFromContext(r.Context())
		idStr := params.ByName("id")
		if idStr == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}*/
		id, err := app.GetIntParam("id", r)
		if err != nil {
			RespondWithError(w, http.StatusBadRequest, "There is no UserID!")
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
			RespondWithError(w, http.StatusBadRequest, "Payload could not be parsed!")
			return
		}
		defer r.Body.Close()

		pw_hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, "An error occurred while hashing the password!")
			return
		}
		user.Password = string(pw_hash)

		id, err := u.Create(&user)
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, "Could not create user!")
			return
		}

		userNew, err := u.GetById(id)
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, "Something went wrong!")
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
			RespondWithError(w, http.StatusBadRequest, "Payload could not be parsed!")
			return
		}
		defer r.Body.Close()

		err := app.CheckUser(user.Id, r)
		if err != nil {
			RespondWithError(w, http.StatusBadRequest, "You are not authorized to perform this operation!")
			return
		}

		pw_hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, "An error occurred while hashing the password!")
			return
		}
		user.Password = string(pw_hash)

		err = u.Update(&user)
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, "An error occurred while updating user!")
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func DeleteUser(repo repository.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		/*params := httprouter.ParamsFromContext(r.Context())
		idStr := params.ByName("id")
		if idStr == "" {
			RespondWithError(w, http.StatusBadRequest, "There is no UserID in parameters!")
			return
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			RespondWithError(w, http.StatusBadRequest, "An error occured while parsing UserID!")
			return
		}*/
		id, err := app.GetIntParam("id", r)
		if err != nil {
			RespondWithError(w, http.StatusBadRequest, "An error occured while getting UserID!")
		}

		err = app.CheckUser(id, r)
		if err != nil {
			RespondWithError(w, http.StatusUnauthorized, "You are not authorized to perform this operation!")
			return
		}

		err = repo.Delete(id)
		if err != nil {
			RespondWithError(w, http.StatusBadRequest, "User could not be deleted!")

			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
