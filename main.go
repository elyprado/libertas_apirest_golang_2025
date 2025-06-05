package main

import (
	"log"
	"net/http"
	"apigolang/routers"
)

func main() {
	router := routers.SetupRouter()
	log.Println("Iniciando servidor em 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

