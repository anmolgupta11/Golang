package mysqldb

import (
	"database/sql"
	"fmt"

	// github.com/go-sql-driver/mysql is needed to connect with mysql db
	_ "github.com/go-sql-driver/mysql"
)

const (
	dbHost = "localhost"
	dbPort = 5432
	dbUser = "Go_Project"
	dbName = "mysql"
)

// OpenDBConnection is used to open db connection
func OpenDBConnection() *sql.DB {
	mysqlInfo := fmt.Sprintf("host=%s	port=%d	user=%s	dbname=%s	sslmode=disable", dbHost, dbPort, dbUser, dbName)
	db, err := sql.Open("mySQL", mysqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connection to database successful!")

	return db
}
