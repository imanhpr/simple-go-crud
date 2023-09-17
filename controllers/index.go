package index

import (
	"encoding/json"
	"fmt"
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
	}
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintln(w, string(marshal))
}
