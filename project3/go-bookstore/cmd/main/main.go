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
	log.Fatal(http.ListenAndServe(":9090", r))
}
