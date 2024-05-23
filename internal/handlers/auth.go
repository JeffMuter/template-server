package handlers

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"

	"template-server/internal/database"
	"template-server/models"
)

var sessions = map[string]string{}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {

	var user models.User

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "error parsing form", http.StatusBadRequest)
		return
	}

	user.Email = r.FormValue("email")
	user.Password = r.FormValue("password")

	if user.Email == "" || user.Password == "" {
		http.Error(w, "Email and password are required", http.StatusBadRequest)
		return
	}

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

	err = setSession(user.Email, w)
	if err != nil {
		http.Error(w, "Failed to set session", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func LoginFormHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "error parsing form", http.StatusBadRequest)
		return
	}

	user.Email = r.FormValue("email")
	user.Password = r.FormValue("password")

	isValidated, err := validateUsernamePassword(user.Email, user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	if isValidated {
		err := setSession(user.Email, w)
		if err != nil {
			http.Error(w, "Failed to set session", http.StatusBadRequest)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		http.Error(w, "Username or password incorrect... Try again.", http.StatusBadRequest)
	}
}

func validateUsernamePassword(email string, password string) (bool, error) {
	// open db connection
	db := database.DatabaseOpen()
	defer db.Close()

	var hashedPassword string
	query := "SELECT password FROM users WHERE email = $1"
	err := db.QueryRow(query, email).Scan(&hashedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, err
		}
		return false, err
	}

	// compare provided password to the password from the db
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false, fmt.Errorf("invalid password")
	}

	return true, nil
}

func genSessionId() (string, error) {
	byteSlice := make([]byte, 32)
	_, err := rand.Read(byteSlice)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(byteSlice), nil
}

func setSession(email string, w http.ResponseWriter) error {
	sessionId, err := genSessionId()
	if err != nil {
		return err
	}
	sessions[sessionId] = email

	http.SetCookie(w, &http.Cookie{
		Name:    "session_id",
		Value:   sessionId,
		Expires: time.Now().Add(24 * time.Hour),
	})
	return nil
}

func getSession(r *http.Request) (string, error) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		return "", err
	}
	userName, exists := sessions[cookie.Value]
	if !exists {
		return "", http.ErrNoCookie
	}
	return userName, nil
}
