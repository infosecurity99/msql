package main

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type Countries struct {
	ID       uuid.UUID
	Name     string
	Currency string
}

func main() {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=admin password=1999 database=news sslmode=disable")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	country := Countries{}

	row := db.QueryRow(`SELECT id,  name, currency FROM  countries  WHERE   id=$1`, "0188f138-84f6-4130-9454-724a0ca19842")

	if err = row.Scan(&country.ID, &country.Name, &country.Currency); err != nil {
		panic(err)
	}

	fmt.Println("Country ", country)
}
