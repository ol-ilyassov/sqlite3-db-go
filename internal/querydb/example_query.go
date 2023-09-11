package querydb

import (
	"database/sql"
	"fmt"
)

// ExampleQuery displays example of db.Query(stmt) function.
// Query() method runs a SELECT query that returns rows and accepts optional arguments as placeholder parameters.
func ExampleQuery(db *sql.DB) error {
	rows, err := db.Query("SELECT * FROM students where score > 8")
	if err != nil {
		return err
	}
	defer rows.Close()

	// iterate through all the records
	for rows.Next() {
		var id int
		var name string
		var score float64
		err = rows.Scan(&id, &name, &score)
		if err != nil {
			return err
		}
		fmt.Printf("%v %v %v\n", id, name, score)
	}
	return nil
}
