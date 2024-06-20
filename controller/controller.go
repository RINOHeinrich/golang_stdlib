package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/RINOHeinrich/golang_stdlib/database"
	"github.com/RINOHeinrich/golang_stdlib/models"
)

func GetAllBooks(w *http.ResponseWriter, r *http.Request) {
	cmd := "SELECT * FROM book"
	var books []models.Book
	stmt, err := database.DatabaseConnection.Prepare(cmd)
	if err != nil {
		fmt.Println("Error preparing statement: ", err)
	}
	bookRows, err := stmt.Query()
	if err != nil {
		fmt.Println("Error querying: ", err)

	}
	defer bookRows.Close()
	for bookRows.Next() {
		var book models.Book
		if err := bookRows.Scan(&book.ID, &book.Title, &book.Author, &book.Published_date); err != nil {
			fmt.Println("Error scanning row: ", err)
			return
		}
		books = append(books, book)
		// Assign the result of append to a variable to avoid the error.
	}
	json.NewEncoder(*w).Encode(books)
}
func GetBook(w *http.ResponseWriter, r *http.Request) {
	book := models.Book{}
	query := r.URL.Query()
	id_str := &query["id"][0]
	id, _ := strconv.Atoi(*id_str)
	cmd := "SELECT * FROM book WHERE id=$1"
	stmt, err := database.DatabaseConnection.Prepare(cmd)
	if err != nil {
		fmt.Println("Error preparing statement: ", err)
		return
	}
	defer stmt.Close()
	bookRow := stmt.QueryRow(id)
	if err := bookRow.Scan(&book.ID, &book.Title, &book.Author, &book.Published_date); err != nil {
		fmt.Println("Error scanning row: ", err)
		return
	}
	json.NewEncoder(*w).Encode(book)
}
func InsertBook(w *http.ResponseWriter, r *http.Request) {
	book := models.Book{}
	json.NewDecoder(r.Body).Decode(&book)
	// Handle POST request
	stmt, err := database.DatabaseConnection.Prepare("INSERT INTO book (title, author, published_date) VALUES ($1, $2, $3)")
	if err != nil {
		fmt.Println("Error preparing statement: ", err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(book.Title, book.Author, book.Published_date)
	if err != nil {
		fmt.Println("Error executing statement: ", err)
		return
	}
	fmt.Fprintf(*w, "Book inserted: %v\n", book)
}

func UpdateBook(w *http.ResponseWriter, r *http.Request) {
	book := models.Book{}
	json.NewDecoder(r.Body).Decode(&book)
	query := r.URL.Query()
	id_str := &query["id"][0]
	id, _ := strconv.Atoi(*id_str)
	if *id_str == "" {
		http.Error(*w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	cmd := "UPDATE book SET title=$1,author=$2,published_date=$3 WHERE id=$4"
	stmt, err := database.DatabaseConnection.Prepare(cmd)
	if err != nil {
		fmt.Println("Error preparing statement: ", err)
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(book.Title, book.Author, book.Published_date, id)
	if err != nil {
		fmt.Println("Error executing statement: ", err)
		return
	}
	fmt.Fprintf(*w, "Book Updated:  %v \n", book)
}
func DeleteBook(w *http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	id_str := &query["id"][0]
	if *id_str == "" {
		http.Error(*w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	id, _ := strconv.Atoi(*id_str)
	cmd := "DELETE FROM book WHERE id=$1"
	stmt, err := database.DatabaseConnection.Prepare(cmd)
	if err != nil {
		fmt.Println("Error preparing statement: ", err)
		return
	}
	defer stmt.Close()
	stmt.Exec(id)
	fmt.Fprintf(*w, "DELETE request received on id %s", *id_str)
}
