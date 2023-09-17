package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/imanhpr/simple-go-crud/helpers"
	"github.com/imanhpr/simple-go-crud/models"
	"github.com/imanhpr/simple-go-crud/routes"
	_ "github.com/mattn/go-sqlite3"
)

const PORT = ":3001"

func createTables(db *sql.DB) {
	_, err := db.Exec(models.CREATE_USER_TABLE_STATMENT)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("User table init successfully.")
}

func main() {
	db := helpers.Database
	createTables(db)
	defer db.Close()
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Mount("/", routes.IndexRouter())
	log.Printf("Server Is Listening On Port %s\n", PORT)
	http.ListenAndServe(PORT, r)
}
