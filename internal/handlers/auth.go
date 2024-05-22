package handlers

import (
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"template-server/internal/database"
	"template-server/models"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {

	var user models.User

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "error parsing form", http.StatusBadRequest)
		return
	}

	user.Email = r.FormValue("email")
	user.Password = r.FormValue("password")

	// hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}
	user.Password = string(hashedPassword)

	// open db connection
	db := database.DatabaseOpen()
	defer db.Close()

	// create sql statement
	query, err := db.Prepare("INSERT INTO users (email, password) VALUES ($1, $2)")
	if err != nil {
		http.Error(w, "Error preparing query", http.StatusInternalServerError)
		return
	}
	defer query.Close()

	_, err = query.Exec(user.Email, user.Password)
	if err != nil {
		http.Error(w, "Error executing query", http.StatusInternalServerError)
		return
	}
	// response successful
	w.WriteHeader(http.StatusCreated)
}

func LoginFormHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	users := models.GetUsers()

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "error parsing form", http.StatusBadRequest)
		return
	}

	user.Email = r.FormValue("email")
	user.Password = r.FormValue("password")

	for _, currentUser := range users {
		if user.Email == currentUser.Email && user.Password == currentUser.Password {

		}
	}
}
