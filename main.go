package main

import (
	"fmt"
	"net/http"

	"github.com/RINOHeinrich/multiserviceauth/controller"
	"github.com/RINOHeinrich/multiserviceauth/database"
)

func main() {
	database.CreateBookTable(database.DatabaseConnection)
	http.HandleFunc("/", bookhandler)
	fmt.Println("Server started on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func bookhandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	fmt.Println(query)
	switch r.Method {
	case "GET":
		// Handle GET request
		if len(query["id"]) == 0 {
			controller.GetAllBooks(&w, r)
			return
		}
		controller.GetBook(&w, r)
	case "POST":
		controller.InsertBook(&w, r)
	case "PUT":
		controller.UpdateBook(&w, r)
	case "DELETE":
		controller.DeleteBook(&w, r)
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
