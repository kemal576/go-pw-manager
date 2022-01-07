package router

import (
	"github.com/julienschmidt/httprouter"
	"github.com/kemal576/go-pw-manager/internal/api"
	"github.com/kemal576/go-pw-manager/repository"
)

type Router struct {
	Router *httprouter.Router
	repos  repository.Database
}

// New ...
func New(rp repository.Database) *Router {
	r := &Router{
		Router: httprouter.New(),
		repos:  rp,
	}
	r.initRoutes()
	return r
}

func (r *Router) initRoutes() {
	// User endpoints
	r.Router.HandlerFunc("GET", "/users", api.IsAuthorized(api.AllUsers(r.repos.Users())))
	r.Router.HandlerFunc("GET", "/users/:id", api.IsAuthorized(api.GetUser(r.repos.Users())))
	r.Router.HandlerFunc("POST", "/users", api.CreateUser(r.repos.Users()))
	r.Router.HandlerFunc("PUT", "/users", api.IsAuthorized(api.UpdateUser(r.repos.Users())))
	r.Router.HandlerFunc("DELETE", "/users/:id", api.IsAuthorized(api.DeleteUser(r.repos.Users())))

	//Login endpoints
	r.Router.HandlerFunc("GET", "/logins", api.IsAuthorized(api.GetLogins(r.repos.Logins())))
	r.Router.HandlerFunc("GET", "/user/:id/logins", api.IsAuthorized(api.GetLoginsByUserId(r.repos.Logins())))
	r.Router.HandlerFunc("GET", "/logins/:id", api.IsAuthorized(api.GetLoginById(r.repos.Logins())))
	r.Router.HandlerFunc("POST", "/logins", api.IsAuthorized(api.CreateLogin(r.repos.Logins())))
	r.Router.HandlerFunc("PUT", "/logins", api.IsAuthorized(api.UpdateLogin(r.repos.Logins())))
	r.Router.HandlerFunc("DELETE", "/logins/:id", api.IsAuthorized(api.DeleteLogin(r.repos.Logins())))

	// Auth endpoints
	r.Router.HandlerFunc("POST", "/signin", api.SignIn(r.repos.Users()))
	//r.Router.HandlerFunc("GET", "/signout", api.Signout())

}
