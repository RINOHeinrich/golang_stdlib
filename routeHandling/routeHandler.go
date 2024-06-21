package requestHandler

import (
	"fmt"
	"net/http"

	"github.com/RINOHeinrich/multiserviceauth/controller"
)

func userHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	fmt.Println(query)
	switch r.Method {
	case "GET":
		// Handle GET request
		if len(query["id"]) == 0 {
			controller.GetAllUsers(&w, r)
			return
		}
		controller.GetUser(&w, r)
	case "POST":
		controller.InsertUser(&w, r)
	case "PUT":
		controller.UpdateUser(&w, r)
	case "DELETE":
		controller.DeleteUser(&w, r)
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
func loginHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	fmt.Println(query)
	switch r.Method {
	case "POST":
		//controller.Login(&w, r)
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
