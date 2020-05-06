package routers

import "github.com/gorilla/mux"

// InitRoutes to call the router mux
func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router = APIRoutes(router)
	router = WebRoutes(router)
	return router
}
