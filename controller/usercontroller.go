package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/RINOHeinrich/multiserviceauth/database"
	"github.com/RINOHeinrich/multiserviceauth/helper"
	"github.com/RINOHeinrich/multiserviceauth/models"
)

func GetAllUsers(w *http.ResponseWriter, r *http.Request) {
	cmd := "SELECT * FROM users"
	var users []models.User
	stmt, err := database.DatabaseConnection.Prepare(cmd)
	if err != nil {
		fmt.Println("Error preparing statement: ", err)
	}
	userRows, err := stmt.Query()
	if err != nil {
		fmt.Println("Error querying: ", err)

	}
	defer userRows.Close()
	for userRows.Next() {
		var user models.User
		if err := userRows.Scan(&user.ID, &user.Username, &user.Password, &user.Email); err != nil {
			fmt.Println("Error scanning row: ", err)
			return
		}
		users = append(users, user)
		// Assign the result of append to a variable to avoid the error.
	}
	json.NewEncoder(*w).Encode(users)
}
func GetUser(w *http.ResponseWriter, r *http.Request) {
	user := models.User{}
	query := r.URL.Query()
	id_str := &query["id"][0]
	id, _ := strconv.Atoi(*id_str)
	cmd := "SELECT * FROM users WHERE id=$1"
	stmt, err := database.DatabaseConnection.Prepare(cmd)
	if err != nil {
		fmt.Println("Error preparing statement: ", err)
		return
	}
	defer stmt.Close()
	userRow := stmt.QueryRow(id)
	if err := userRow.Scan(&user.ID, &user.Username, &user.Password, &user.Email); err != nil {
		fmt.Println("Error scanning row: ", err)
		return
	}
	json.NewEncoder(*w).Encode(user)
}
func InsertUser(w *http.ResponseWriter, r *http.Request) {
	user := models.User{}
	json.NewDecoder(r.Body).Decode(&user)
	// Handle POST request
	user.Password = helper.HashPassword(user.Password)
	stmt, err := database.DatabaseConnection.Prepare("INSERT INTO users (username, password, email) VALUES ($1, $2, $3)")
	if err != nil {
		fmt.Println("Error preparing statement: ", err)
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.Username, user.Password, user.Email)
	if err != nil {
		fmt.Println("Error executing statement: ", err)
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
	id, _ := strconv.Atoi(*id_str)
	cmd := "UPDATE users SET username=$1, password=$2, email=$3 WHERE id=$4"
	stmt, err := database.DatabaseConnection.Prepare(cmd)
	if err != nil {
		fmt.Println("Error preparing statement: ", err)
		return
	}
	defer stmt.Close()
	stmt.Exec(user.Username, user.Password, user.Email, id)
	fmt.Fprintf(*w, "PUT request received on id %s", *id_str)
}
func DeleteUser(w *http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	id_str := &query["id"][0]
	if *id_str == "" {
		http.Error(*w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	id, _ := strconv.Atoi(*id_str)
	cmd := "DELETE FROM users WHERE id=$1"
	stmt, err := database.DatabaseConnection.Prepare(cmd)
	if err != nil {
		fmt.Println("Error preparing statement: ", err)
		return
	}
	defer stmt.Close()
	stmt.Exec(id)
	fmt.Fprintf(*w, "DELETE request received on id %s", *id_str)
}
