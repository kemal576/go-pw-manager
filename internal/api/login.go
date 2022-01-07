package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/kemal576/go-pw-manager/internal/app"
	"github.com/kemal576/go-pw-manager/models"
	"github.com/kemal576/go-pw-manager/repository"
)

func GetLogins(repo repository.LoginRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logins, err := app.GetAllLogins(repo)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		RespondWithJSON(w, http.StatusOK, logins)
	}
}

func GetLoginsByUserId(repo repository.LoginRepository) http.HandlerFunc {
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

		logins, err := app.GetLoginsByUserId(id, repo)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			println(err.Error())
			return
		}
		RespondWithJSON(w, http.StatusOK, logins)
	}
}

func GetLoginById(repo repository.LoginRepository) http.HandlerFunc {
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
}

func CreateLogin(repo repository.LoginRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var login models.LoginDTO

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&login); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		id, err := app.CreateLogin(login, repo)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		RespondWithJSON(w, http.StatusCreated, id)
	}
}

func UpdateLogin(repo repository.LoginRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var login models.Login

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&login); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		id, err := app.UpdateLogin(login, repo)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		RespondWithJSON(w, http.StatusCreated, id)
	}
}

func DeleteLogin(repo repository.LoginRepository) http.HandlerFunc {
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

		err = repo.Delete(id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		RespondWithJSON(w, http.StatusOK, id)
	}
}
