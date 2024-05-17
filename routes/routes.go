package routes

import (
	"net/http"
	"template-server/handlers"
)

func Router() *http.ServeMux {

	mux := http.NewServeMux()

	// Register handlers
	mux.HandleFunc("/post/", handlers.PostHandler)
	mux.HandleFunc("/login", handlers.LoginHandler)
	mux.HandleFunc("/", handlers.HomeHandler)

	return mux

}
