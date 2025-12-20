package main

import (
	router "api-service/internal/http"
	"log"
	"net/http"
)

func main() {
	router := router.NewRouter()
	log.Println("Service Starting on port 2222")
	err := http.ListenAndServe(":2222", router)
	if err != nil {
		log.Fatal(err)
	}
}
