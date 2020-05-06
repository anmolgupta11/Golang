package handlers

import (
	"Go_Project/mysqldb"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// ResponseError struct
type ResponseError struct {
	Status  int    `json."status"`
	Message string `json."message"`
}

// GetBook function to get a book
func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	idStr := params["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		APIError(w, 400, "Bad Request. Invalid ID")
		return
	}

	book, err := mysqldb.GetOneBook(id)
	if err != nil {
		APIError(w, 404, err.Error())
		return
	}

	json.NewEncoder(w).Encode(book)
}

// GetBooks function to get all books
func GetBooks(w http.ResponseWriter, r *http.Request) {
	books, err := mysqldb.GetAllBooks()
	if err != nil {
		APIError(w, 200, err.Error())
	}

	json.NewEncoder(w).Encode(books)
}

// APIError to check for responses
func APIError(w http.ResponseWriter, status int, m string) {
	e := ResponseError{
		Status:  status,
		Message: m,
	}
	json.NewEncoder(w).Encode(e)
}
