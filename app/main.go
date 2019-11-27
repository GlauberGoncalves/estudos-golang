package main

import (
	"estudos-golang/app/routes"
	"log"
	"net/http"
)

func main() {

	routes.LoadRoutes()

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Println(err)
	}

	log.Println("online on localhost:3000")
}
