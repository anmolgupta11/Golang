package mysqldb

import "database/sql"

// SQLInsert interface handles inserting elements into the respective tables
type SQLInsert interface {
	insertRow(*sql.DB) bool
}

// Insert is used to insert row
func Insert(si SQLInsert) bool {
	db := OpenDBConnection()
	ok := si.insertRow(db)
	if !ok {
		return false
	}
	return true
}
