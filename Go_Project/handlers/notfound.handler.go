package handlers

import (
	"Go_Project/templates"
	"net/http"
)

// NotFoundHandler to check for 404 error if no handler found
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	t := &templates.Page{Title: "Not Found"}
	templates.RenderTemplate(w, "templates/404", t)
}
