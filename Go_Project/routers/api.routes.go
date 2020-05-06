package routers

import (
	"Go_Project/handlers"

	"github.com/gorilla/mux"
)

// APIRoutes to return api routes
func APIRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/book/get/{id}", handlers.GetBook)
	router.HandleFunc("/books/get", handlers.GetBooks)
	return router
}
