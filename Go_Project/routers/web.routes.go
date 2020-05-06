package routers

import (
	"Go_Project/handlers"

	"github.com/gorilla/mux"
)

// WebRoutes to handle route requests from webserver
func WebRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("404", handlers.NotFoundHandler)
	router.HandleFunc("/read-file/", handlers.FileReadHandler)
	router.HandleFunc("/books/", handlers.BooksViewHandler)
	router.HandleFunc("/get-params/", handlers.GetParamsViewHandler)
	router.HandleFunc("/elements/", handlers.ElementsViewHandler)
	router.HandleFunc("/", handlers.HomeViewHandler)

	return router
}
