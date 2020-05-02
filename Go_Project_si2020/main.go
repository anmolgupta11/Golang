package main

import (
	"log"
	"net/http"

	"Go_Project_si2020/config"
	"Go_Project_si2020/routers"
)

func main() {
	router := routers.RouterMux()

	config.ConnectCache()

	defer config.Cache.Close()

	config.PingCache()

	err := http.ListenAndServe(":8081", router)

	if err != nil {
		log.Fatal(err)

	}

}
