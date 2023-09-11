package main

import (
	"database/sql"
	"log"

	"sqlite3-db/internal/querydb"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Make connection:
	db, err := sql.Open("sqlite3", "./sqlite3-app.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = querydb.PrepareDB(db)
	if err != nil {
		log.Fatal(err)
	}

	err = querydb.CRUD(db)
	if err != nil {
		log.Fatal(err)
	}
}
