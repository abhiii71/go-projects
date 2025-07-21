package main

import (
	"log"
	"net/http"

	"github.com/abhiii71/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	routes.RegisterBookStoreRoute(r)
	http.Handle("/", r)
	port := ":9090"
	log.Println("Server started on port", port)
	log.Fatal(http.ListenAndServe(port, r))
}
