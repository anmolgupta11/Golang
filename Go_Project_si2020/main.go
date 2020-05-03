package main

import (
	"net/http"

	"Go_Project_si2020/config"
	"Go_Project_si2020/routers"
	"Go_Project_si2020/settings"

	"github.com/codegangsta/negroni"
)

func main() {

	settings.Init()
	router := routers.InitRoutes()

	config.ConnectDatabase()
	config.ConnectCache()

	defer config.Cache.Close()

	config.PingCache()

	n := negroni.Classic()
	n.UseHandler(router)

	http.ListenAndServe(":8081", n)

}
