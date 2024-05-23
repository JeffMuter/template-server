package routes

import (
	"net/http"

	"template-server/internal/handlers"
	"template-server/internal/middleware"
)

func Router() *http.ServeMux {

	mux := http.NewServeMux()

	// Register handlers
	mux.Handle("/", middleware.AuthMiddleware(http.HandlerFunc(handlers.ServeHomePage)))
	// mux.HandleFunc("/", handlers.ServeHomePage)
	mux.HandleFunc("/post/", handlers.ServePostPage)

	mux.HandleFunc("GET /login", handlers.ServeLoginPage)
	mux.HandleFunc("POST /login", handlers.LoginFormHandler)

	mux.HandleFunc("GET /registeruser", handlers.ServeRegistrationPage)
	mux.HandleFunc("POST /registeruser", handlers.RegisterHandler)

	return mux

}
