package user

import (
	"fmt"
	"log"
	"template-server/database"

	_ "github.com/lib/pq"
)

type User struct {
	Id    int
	Name  string
	Email string
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
		err := rows.Scan(&user.Id, &user.Name, &user.Email)
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
		fmt.Printf("ID: %d, Name: %s, Email: %s\n", user.Id, user.Name, user.Email)
	}

	return users
}
