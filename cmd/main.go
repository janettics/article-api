package main

import (
	"log"
	"net/http"

	"../handlers"
	"../svc"
)

func main() {
	rParams := handlers.RouterParams{
		ArticleService: svc.ArticleService(),
	}

	r := handlers.SetupRouter(rParams)
	http.Handle("/", r)

	log.Print("Listening on port :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
