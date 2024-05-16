package user

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type User struct {
	Id    int
	Name  string
	Email string
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Everythingx12"
	dbname   = "trial"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
