package main

import (
	"log"
	"net/http"
	"store-api/config"
	"store-api/routes"
)

func main() {
	config.LoadConfig()
	router := routes.SetupRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
