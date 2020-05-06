package mysqldb

import (
	"Go_Project/redis"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"
)

// Book struct
type Book struct {
	redis.RedisBook
}

// GetAllBooks func to get all the books from db
func GetAllBooks() ([]Book, error) {
	db := OpenDBConnection()
	selectStatement := `SELECT * FROM books ORDER BY id DESC`
	rows, err := db.Query(selectStatement)
	if err != nil {
		fmt.Println(err.Error())
		return []Book{}, err
	}
	var books []Book
	bookObj := Book{}

	for rows.Next() {
		var (
			id     int
			name   string
			author string
			pages  int
		)
		err = rows.Scan(&id, &name, &author, &pages)
		if err != nil {
			fmt.Println(err.Error())
		}
		bookObj.ID = id
		bookObj.Name = name
		bookObj.Author = author
		bookObj.Pages = pages
	}

	defer db.Close()
	return books, err
}

// GetOneBook to get one book acc to ID from db
func GetOneBook(id int) (Book, error) {
	var (
		name   string
		author string
		pages  int
	)

	// checking in redis first

	val, err := redis.GetBookCache("book_" + strconv.Itoa(id))
	book := Book{val}
	if err != nil {
		log.Println(err.Error())
	} else {
		return book, nil
	}

	// if nothing found in cache look in database

	sqlStatement := `SELECT * FROM books WHERE id=$1`
	db := OpenDBConnection()
	row := db.QueryRow(sqlStatement, id)
	switch err := row.Scan(&id, &name, &author, &pages); err {
	case sql.ErrNoRows:
		return Book{}, errors.New("Book does not exist")
	case nil:
		book.ID = id
		book.Name = name
		book.Author = author
		book.Pages = pages

		// As the book was not found in cache, therefore set it
		err = book.SetBookCache("book_" + strconv.Itoa(id))
		if err != nil {
			log.Println(err.Error())
		}
		return book, nil

	default:
		return Book{}, errors.New("An unknown error occurred")
	}
}

// InsertRow Adding a single book object to database
func (b Book) insertRow(db *sql.DB) bool {
	insertStatement := `INSERT INTO books (name, author, pages) VALUES ($1, $2, $3)`
	_, err := db.Exec(insertStatement, b.Name, b.Author, b.Pages)
	if err != nil {
		return false
	}
	return true
}
