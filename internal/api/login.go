package api

import (
	"encoding/json"
	"net/http"

	"github.com/kemal576/go-pw-manager/internal/app"
	"github.com/kemal576/go-pw-manager/models"
	"github.com/kemal576/go-pw-manager/repository"
	"github.com/kemal576/go-pw-manager/utils"
)

func GetLogins(repo repository.LoginRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logins, err := app.GetAllLogins(repo)
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, "Something went wrong!")
			return
		}

		if len(logins) == 0 {
			RespondWithError(w, http.StatusNotFound, "No logins found!")
			return
		}
		RespondWithJSON(w, http.StatusOK, logins)
	}
}

func GetLoginsByUserId(repo repository.LoginRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := utils.GetIntParam("id", r)
		if err != nil {
			RespondWithError(w, http.StatusBadRequest, "There is no UserID!")
			return
		}

		check := app.CheckUser(id, r)
		if check != true {
			RespondWithError(w, http.StatusUnauthorized, "You are not authorized to perform this operation!")
			return
		}

		logins, err := app.GetLoginsByUserId(id, repo)
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, "Something went wrong!")
			return
		}

		if len(logins) == 0 {
			RespondWithError(w, http.StatusNotFound, "No logins found!")
			return
		}
		RespondWithJSON(w, http.StatusOK, logins)
	}
}

func GetLoginById(repo repository.LoginRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := utils.GetIntParam("id", r)
		if err != nil {
			RespondWithError(w, http.StatusBadRequest, "There is no LoginID!")
		}

		login, err := repo.GetById(id)
		if err != nil || login.Id == 0 {
			RespondWithError(w, http.StatusInternalServerError, "Login not found!")
			return
		}
		RespondWithJSON(w, http.StatusOK, login)
	}
}

func CreateLogin(repo repository.LoginRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var login models.LoginDTO

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&login); err != nil {
			RespondWithError(w, http.StatusBadRequest, "Payload could not be parsed!")
			return
		}
		defer r.Body.Close()

		id, err := app.CreateLogin(login, repo)
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, "Login could not create!")
			return
		}

		newLogin, err := repo.GetById(id)
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, "Something went wrong!")
			return
		}

		RespondWithJSON(w, http.StatusCreated, newLogin)
	}
}

func UpdateLogin(repo repository.LoginRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var login models.Login

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&login); err != nil {
			RespondWithError(w, http.StatusUnprocessableEntity, "Payload could not be parsed!")
			return
		}
		defer r.Body.Close()

		check := app.CheckUser(login.UserId, r)
		if check != true {
			RespondWithError(w, http.StatusUnauthorized, "You are not authorized to perform this operation!")
			return
		}

		_, err := app.UpdateLogin(login, repo)
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, "Login could not updated!")
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func DeleteLogin(repo repository.LoginRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := utils.GetIntParam("id", r)
		if err != nil {
			RespondWithError(w, http.StatusBadRequest, "There is no LoginID!")
		}

		login, err := repo.GetById(id)
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, "Login not found!")
		}

		check := app.CheckUser(login.UserId, r)
		if check != true {
			RespondWithError(w, http.StatusUnauthorized, "You are not authorized to perform this operation!")
			return
		}

		err = repo.Delete(id)
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, "Login could not deleted!")
			return
		}
		RespondWithJSON(w, http.StatusOK, id)
	}
}

/*
func GetLoginByUrl(repo repository.LoginRepository) http.HandlerFunc {
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
		login, err := repo.GetById(id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		RespondWithJSON(w, http.StatusOK, login)
	}
}*/
