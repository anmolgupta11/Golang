package handlers

import (
	"Go_Project/mysqldb"
	"Go_Project/redis"
	"Go_Project/templates"
	"net/http"
	"strconv"
)

// BooksViewHandler func for viewing books
func BooksViewHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			templates.BasicHTTPRender(w, "Failed to Post data.")
			return
		}
		name := r.FormValue("name")
		author := r.FormValue("author")
		pageStr := r.FormValue("pages")

		pages, err := strconv.Atoi(pageStr)
		if err != nil {
			templates.BasicHTTPRender(w, "Pages must be a valid integer.")
		}
		bookObj := mysqldb.Book{redis.RedisBook{
			Name:   name,
			Author: author,
			Pages:  pages,
		},
		}
		ok := mysqldb.Insert(bookObj)
		if !ok {
			templates.BasicHTTPRender(w, "Failed to add the book to database.")
		}
	}
	books, err := mysqldb.GetAllBooks()
	if err != nil {
		templates.BasicHTTPRender(w, "Error reading from daabase.")
		return
	}
	templates.RenderTemplate(w, "templates/books", &templates.Page{Title: "Books", Books: books})
}
