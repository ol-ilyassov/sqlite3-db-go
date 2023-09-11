package querydb

import (
	"database/sql"
	"log"
)

// Testconnection checks database connection and returns database version.
func TestConnection(db *sql.DB) error {
	var err error
	var version string

	err = db.QueryRow("SELECT SQLITE_VERSION()").Scan(&version)
	if err != nil {
		return err
	}

	log.Println("SQLite3 version:", version)
	return nil
}
