package handlers

import (
	"Go_Project/templates"
	"io/ioutil"
	"net/http"
)

// FileReadHandler to read the files
func FileReadHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/read-file/"):]
	p, err := loadPageFromFile(title)
	if err != nil {
		templates.BasicHTTPRender(w, "Error opening that file. Expected path: /read-file/{{filename}}  [please do not use file extension]")
		return
	}
	templates.RenderTemplate(w, "templates/read-file", p)
}

// loadPageFromFile to load the pages from file
func loadPageFromFile(title string) (*templates.Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &templates.Page{Title: title, Body: string(body)}, nil
}
