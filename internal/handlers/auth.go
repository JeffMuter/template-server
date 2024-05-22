package handlers

import (
	"net/http"

	"golang.org/x/crypto/bcrypt"

	user "template-server/models"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {

	var user user.User

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "error parsing form", http.StatusBadRequest)
		return
	}

	user.Email = r.FormValue("email")
	user.Password = r.FormValue("password")

	//hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}
	user.Password = string(hashedPassword)

	//add user to db

}
