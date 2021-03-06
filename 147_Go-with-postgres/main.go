package main

import (
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "postgres://gotraining:password@localhost/bookstore?sslmode=disable")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("You are connected")
}