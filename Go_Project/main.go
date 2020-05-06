package main

import (
	"Go_Project/routers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r = routers.InitRoutes()
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
