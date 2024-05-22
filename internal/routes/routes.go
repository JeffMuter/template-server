package routes

import (
	"net/http"
	"template-server/internal/handlers"
)

func Router() *http.ServeMux {

	mux := http.NewServeMux()

	// Register handlers
	mux.HandleFunc("/post/", handlers.PostHandler)
	mux.HandleFunc("GET /login", handlers.ServeLoginForm)
	mux.HandleFunc("POST /login", handlers.LoginFormHandler)

	mux.HandleFunc("/", handlers.HomeHandler)
	mux.HandleFunc("GET /registeruser", handlers.ServeRegistrationForm)
	mux.HandleFunc("POST /registeruser", handlers.RegisterHandler)

	return mux

}
