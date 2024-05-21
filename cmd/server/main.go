package main

import (
	"fmt"
	"log"
	"net/http"

	"template-server/internal/routes"
	user "template-server/models"
)

func main() {

	fmt.Printf("All users: %v", user.GetUsers())
	mux := routes.Router()

	log.Printf("Started server on 8080:\nhttp://localhost:8080")

	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		log.Fatal("Error starting server: ", err)
	}

}
