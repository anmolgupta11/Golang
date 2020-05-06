package mysqldb

import (
	"database/sql"
	"fmt"
)

// Element declares the schema for the table "elements"
type Element struct {
	ID   int
	Text string
}

// GetAllElements fetches all rows from the "elements" table and returns a
// slice []Element
func GetAllElements() ([]Element, error) {
	db := OpenDBConnection()
	selectStatement := `SELECT * FROM elements ORDER BY id DESC`
	rows, err := db.Query(selectStatement)
	if err != nil {
		fmt.Println(err.Error())
		return []Element{}, err
	}
	elements := []Element{}
	elementObj := Element{}

	for rows.Next() {
		var (
			id      int
			element string
		)
		err = rows.Scan(&id, &element)
		if err != nil {
			fmt.Println(err.Error())
		}
		elementObj.ID = id
		elementObj.Text = element
		elements = append(elements, elementObj)
	}
	defer db.Close()
	return elements, err
}

// Add a single Element object to the database
func (e Element) insertRow(db *sql.DB) bool {
	insertStatement := `INSERT INTO elements ( element ) VALUES ($1)`
	_, err := db.Exec(insertStatement, e.Text)
	if err != nil {
		return false
	}
	return true
}
