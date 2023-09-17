package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/imanhpr/simple-go-crud/routes"
)

const PORT = ":3001"

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Mount("/", routes.IndexRouter())
	log.Printf("Server Is Listening On Port %s\n", PORT)
	http.ListenAndServe(PORT, r)
}
