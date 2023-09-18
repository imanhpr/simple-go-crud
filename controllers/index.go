package index

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/imanhpr/simple-go-crud/models"
)

func IndexController(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to simple-go-crud")
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users := models.UserRepoInstance.GetAll()
	marshal, err := json.Marshal(users)
	if err != nil {
		w.Header().Add("Content-Type", "text/plain")
		w.WriteHeader(500)
		fmt.Fprintln(w, "something went wrong")
		return
	}
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintln(w, string(marshal))
}

func PostNewUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintln(w, "Data validation faild")
		return
	}
	err = json.Unmarshal(body, &user)
	if err != nil {
		w.WriteHeader(422)
		fmt.Fprintln(w, "faild to deserialize")
		return
	}
	newUser := models.UserRepoInstance.InsertOne(user)
	marshal, err := json.Marshal(newUser)
	if err != nil {
		w.WriteHeader(500)
		fmt.Println("new User insert failed")
		return
	}
	fmt.Fprintln(w, string(marshal))
}
