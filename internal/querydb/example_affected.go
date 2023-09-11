package querydb

import (
	"database/sql"
	"fmt"
)

// ExampleRowsAffected displays example of rows.RowsAffected() function.
// RowsAffected() retuns number of rows that have been affected by an UPDATE, INSERT, or DELETE statement.
func ExampleRowsAffected(db *sql.DB) error {
	query := `
		INSERT INTO students(name, score) VALUES('Kelly',8.0);
		INSERT INTO students(name, score) VALUES('Kai',9.5);
		INSERT INTO students(name, score) VALUES('Ben',9.0);
		INSERT INTO students(name, score) VALUES('Bin',7.5);
	`
	res, err := db.Exec(query)
	if err != nil {
		return err
	}

	numOfRow, err := res.RowsAffected()
	if err != nil {
		return err
	}
	fmt.Printf("The statement has affected %d rows\n", numOfRow)

	return nil
}
