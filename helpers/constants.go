package helpers

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func connectToDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", "db.sqlite3")
	if err != nil {
		log.Println(err)
		log.Fatal("Something went wrong wrong ")
	}
	err = db.Ping()
	if err != nil {
		log.Println(err)
		log.Fatal("Can not ping with database")
	}
	return db
}

var Database = connectToDatabase()
