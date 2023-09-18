package routes

import (
	"github.com/go-chi/chi/v5"
	index "github.com/imanhpr/simple-go-crud/controllers"
)

func IndexRouter() chi.Router {
	r := chi.NewRouter()
	r.Get("/", index.IndexController)
	r.Get("/user", index.GetAllUsers)
	r.Post("/user", index.PostNewUser)
	return r
}
