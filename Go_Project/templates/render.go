package templates

import (
	"Go_Project/mysqldb"
	"fmt"
	"html/template"
	"net/http"
)

// Page struct to define parameters
type Page struct {
	Title    string
	Body     string
	Elements []mysqldb.Element
	Books    []mysqldb.Book
}

// BasicHTTPRender to write simple text based HTTP response
func BasicHTTPRender(w http.ResponseWriter, message string) {
	fmt.Fprintf(w, "%s", message)
	return
}

// RenderTemplate takes page and template as paramters
func RenderTemplate(w http.ResponseWriter, v string, p *Page) {
	t, err := template.ParseFiles(v + ".html")
	if err != nil {
		BasicHTTPRender(w, "Failed to parse the template.")
	}
	err = t.Execute(w, p)
	if err != nil {
		BasicHTTPRender(w, "Failed to parse the template.")
	}
	return
}
