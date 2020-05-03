package config

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

/*Create mysql connection*/
func ConnectDatabase() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8081)/test")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Prepare statement for inserting data
	stmtIns, err := db.Prepare("INSERT INTO test VALUES( ?, ? )") // ? is the variable placeholder
	if err != nil {
		panic(err.Error())
	}
	defer stmtIns.Close() // Close the statement when we leave main() or the program terminates

	// Prepare statement for reading data
	stmtOut, err := db.Prepare("SELECT name FROM test WHERE id > ?") // ? is the variable placeholder
	if err != nil {
		panic(err.Error())
	}
	defer stmtOut.Close() // Close the statement when we leave main() or the program terminates

	// Insert square numbers for 0-24 in the database
	for i := 0; i < 10; i++ {
		_, err = stmtIns.Exec(i, strconv.Itoa(i*100)) // Insert tuples (i, string of i*100)
		if err != nil {
			panic(err.Error())
		}
	}

	// Query the rows that has id more than 5
	rows, err := stmtOut.Query(5)
	if err != nil {
		panic(err.Error())
	}

	// Process each rows accordingly
	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("name is %s\n", name)
	}
}
