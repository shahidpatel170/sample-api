package main

import (
	"log"
	"net/http"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

var middlewares = []Middleware{
	tokenAuthMiddleware,
}

func main() {
	var handler http.HandlerFunc = handleClientProfile
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}
	log.Println("building a RESTful API...")
	http.HandleFunc("/user/profile/{clientId}", handler)

	log.Println("starting server...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
