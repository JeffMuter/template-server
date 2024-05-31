package main

import (
	"log"
	"net/http"

	"template-server/internal/routes"
)

func main() {

	mux := routes.Router()

	log.Printf("\nhttp://localhost:8080")

	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		log.Fatal("Error starting server: ", err)
	}

}
