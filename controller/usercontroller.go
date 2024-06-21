package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/RINOHeinrich/multiserviceauth/database"
	"github.com/RINOHeinrich/multiserviceauth/helper"
	"github.com/RINOHeinrich/multiserviceauth/models"
)

var db database.Postgres

func GetAllUsers(w *http.ResponseWriter, r *http.Request) {

	users, err := database.FindAll(&db)
	if err != nil {
		fmt.Println("Error getting users: ", err)
		return
	}
	json.NewEncoder(*w).Encode(users)
}
func GetUser(w *http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	id_str := &query["id"][0]
	if *id_str == "" {
		http.Error(*w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	user, err := database.Find(&db, *id_str)
	if err != nil {
		fmt.Println("Error getting user: ", err)
		return
	}
	json.NewEncoder(*w).Encode(user)
}
func InsertUser(w *http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	json.NewDecoder(r.Body).Decode(user)
	// Handle POST request
	user.Password = helper.HashPassword(user.Password)
	err := database.Insert(&db, user)
	if err != nil {
		fmt.Println("Error inserting user: ", err)
		return
	}
	fmt.Fprintf(*w, "User inserted: %v\n", user)
}

func UpdateUser(w *http.ResponseWriter, r *http.Request) {
	user := models.User{}
	json.NewDecoder(r.Body).Decode(&user)
	user.Password = helper.HashPassword(user.Password)
	query := r.URL.Query()
	id_str := &query["id"][0]
	if *id_str == "" {
		http.Error(*w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	err := database.Update(&db, *id_str, &user)
	if err != nil {
		fmt.Println("Error updating user: ", err)
		return
	}
	fmt.Fprintf(*w, "PUT request received on id %s", *id_str)
}
func DeleteUser(w *http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	id_str := &query["id"][0]
	if *id_str == "" {
		http.Error(*w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	err := database.Delete(&db, *id_str)
	if err != nil {
		fmt.Println("Error deleting user: ", err)
		return
	}
	fmt.Fprintf(*w, "DELETE request received on id %s", *id_str)
}
