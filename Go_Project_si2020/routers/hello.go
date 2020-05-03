package routers

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func SetHelloRoutes(router *mux.Router) *mux.Router {
	router.Handle(
		"/test/Hello",
		negroni.New(
			negroni.HandlerFunc(controllers.HelloController),
		)).Methods("GET")

	return router
}
