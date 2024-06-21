package controller

import (
	"encoding/json"
	"net/http"

	"github.com/RINOHeinrich/multiserviceauth/models"
)

func Login(w *http.ResponseWriter, r *http.Request) {
	// Handle POST request
	user := models.UserLogin{}
	json.NewDecoder(r.Body).Decode(&user)
	// Check if user exists in the database usin

}
