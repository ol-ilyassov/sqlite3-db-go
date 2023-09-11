package querydb

import (
	"database/sql"
	"log"

	"sqlite3-db/internal/model"
)

// CRUD applies database query operations.
func CRUD(db *sql.DB) error {
	var query string
	var err error

	var tempBook model.Book

	// * ==========================================
	// * SELECT
	// * ==========================================
	query = `
		SELECT id, name, author 
		FROM books
	`
	rows, _ := db.Query(query)

	for rows.Next() {
		err = rows.Scan(&tempBook.ID, &tempBook.Name, &tempBook.Author)
		if err != nil {
			return err
		}
		log.Printf("id:%d, book:%s, author:%s\n", tempBook.ID,
			tempBook.Name, tempBook.Author)
	}

	// * ==========================================
	// * UPDATE
	// * ==========================================
	query = `
		UPDATE books 
		SET name=? 
		WHERE id=?
	`
	statement, _ := db.Prepare(query)
	_, err = statement.Exec("A Tale of Three Cities", 1)
	if err != nil {
		return err
	}
	log.Println("Successfully updated the book in database!")

	// * ==========================================
	// * SELECT after UPDATE
	// * ==========================================
	query = `
		SELECT id, name, author 
		FROM books
	`
	rows, _ = db.Query(query)
	for rows.Next() {
		err = rows.Scan(&tempBook.ID, &tempBook.Name, &tempBook.Author)
		if err != nil {
			return err
		}
		log.Printf("id:%d, book:%s, author:%s\n", tempBook.ID,
			tempBook.Name, tempBook.Author)
	}

	// * ==========================================
	// * DELETE
	// * ==========================================
	query = `
		DELETE FROM books 
		WHERE id=?
	`
	statement, _ = db.Prepare(query)
	_, err = statement.Exec(1)
	if err != nil {
		return err
	}
	log.Println("Successfully deleted the book in database!")
	// * ==========================================

	return nil
}
