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
	r.Router.HandlerFunc("POST", "/users", api.Create(r.repos.Users()))

	// Auth endpoints
	r.Router.HandlerFunc("POST", "/signin", api.SignIn(r.repos.Users()))
	r.Router.HandlerFunc("GET", "/signout", api.Signout())
}
