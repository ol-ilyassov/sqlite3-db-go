package querydb

import (
	"database/sql"
	"fmt"
)

// ExampleExec displays example of db.Exec(stmt) function.
// Exec() method executes a query without returning any rows.
func ExampleExec(db *sql.DB) error {
	var err error

	stmts := `
	DROP TABLE IF EXISTS students;
	CREATE TABLE students(id INTEGER PRIMARY KEY, name TEXT, score REAL);
	INSERT INTO students(name, score) VALUES('Anna',8.5);
	INSERT INTO students(name, score) VALUES('Bob',7.5);
	INSERT INTO students(name, score) VALUES('Claire',9.5);
	INSERT INTO students(name, score) VALUES('Charlie',6.5);
	INSERT INTO students(name, score) VALUES('Daniel',8.0);
	INSERT INTO students(name, score) VALUES('Hellen',7.0);
	INSERT INTO students(name, score) VALUES('Hummer',7.5);
	INSERT INTO students(name, score) VALUES('John',10);
`

	// run the query
	_, err = db.Exec(stmts)
	if err != nil {
		return err
	}
	fmt.Println("table created")

	return nil
}
