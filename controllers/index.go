package index

import (
	"fmt"
	"net/http"
)

func IndexController(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to simple-go-crud")
}
