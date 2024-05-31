package models

import (
	"database/sql"
	"fmt"
	"log"
	"template-server/internal/database"

	_ "github.com/lib/pq"
)

type User struct {
	Id       int
	Email    string
	Password string
}

func GetUserByEmail(email string) (User, error) {
	var user User
	db := database.DatabaseOpen()
	defer db.Close()

	query := `SELECT Id, email, password FROM users WHERE email = $1`
	row := db.QueryRow(query, email)
	err := row.Scan(&user.Id, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Scanned db for a user via email, found none...")
			return user, err
		} else {
			log.Fatal(err)
			return user, err
		}
	}

	return user, nil

}

func GetUsers() []User {
	query := `SELECT * FROM users`

	db := database.DatabaseOpen()
	defer db.Close()

	rows, err := db.Query(query)
	if err != nil {
		log.Fatalf("Error executing query: %v", err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Email)
		if err != nil {
			log.Fatalf("Error scanning row: %v", err)
		}
		users = append(users, user)
	}

	err = rows.Err()
	if err != nil {
		log.Fatalf("Error with rows: %v", err)
	}

	for _, user := range users {
		fmt.Printf("ID: %d, Email: %s\n", user.Id, user.Email)
	}

	return users
}
