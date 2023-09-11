package querydb

import (
	"database/sql"
	"log"
)

// PrepareDB setups "books" table.
func PrepareDB(db *sql.DB) error {
	// * ==========================================
	// * CREATE TABLE
	// * ==========================================
	query := `
		CREATE TABLE 
		IF NOT EXISTS books (
			id INTEGER PRIMARY KEY, 
			isbn INTEGER, 
			author VARCHAR(64), 
			name VARCHAR(64) NULL
		)`
	statement, err := db.Prepare(query)
	if err != nil {
		return nil
	}
	_, err = statement.Exec()
	if err != nil {
		return nil
	}
	log.Println("Successfully created table books!")

	// * ==========================================
	// * INSERT INTO
	// * ==========================================
	query = `
		INSERT INTO books(name, author, isbn) 
		VALUES (?, ?, ?)
	`
	statement, _ = db.Prepare(query)
	_, err = statement.Exec("A Tale of Two Cities", "Charles Dickens", 140430547)
	if err != nil {
		return nil
	}
	log.Println("Inserted the book into database!")
	// * ==========================================

	return nil
}
