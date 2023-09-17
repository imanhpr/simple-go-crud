package models

import (
	"database/sql"
	"log"

	"github.com/imanhpr/simple-go-crud/helpers"
)

const CREATE_USER_TABLE_STATMENT = `
	CREATE TABLE IF NOT EXISTS users (
    	id INTEGER PRIMARY KEY AUTOINCREMENT,
    	name VARCHAR(255)
	);
`

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func NewUser(name string) User {
	return User{Name: name}
}
func NewUserWithId(name string, id int) User {
	return User{Id: id, Name: name}
}

type UserRepo struct {
	Db *sql.DB
}

func (ur UserRepo) InsertOne(user User) {
	const insert = "INSERT INTO users (name) VALUES (?)"
	res, err := ur.Db.Exec(insert, user.Name)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res)
}

func (ur UserRepo) GetAll() []User {
	query := "SELECT * FROM Users"
	rows, err := ur.Db.Query(query)
	if err != nil {
		log.Println(err)
		log.Fatal("Users not found")
	}
	defer rows.Close()
	users := make([]User, 0, 10)
	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Name)
		if err != nil {
			log.Println("I can't  insert this user", err)
			continue
		}
		users = append(users, user)
	}
	return users
}

var UserRepoInstance = UserRepo{Db: helpers.Database}
