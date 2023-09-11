package querydb

import (
	"database/sql"
	"fmt"
)

// ExamplePrepare displays example of db.Prepare(stmt) function.
// Prepare() method setups placeholder values in statement and improves quering database security.
func ExamplePrepare(db *sql.DB) error {
	query := `
		SELECT * 
		FROM students 
		WHERE score > ? and score < ?
	`
	preState, err := db.Prepare(query)
	if err != nil {
		return nil
	}
	minQueryScore := 7
	maxQueryScore := 9

	rows, err := preState.Query(minQueryScore, maxQueryScore)
	if err != nil {
		return err
	}
	defer rows.Close()

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
