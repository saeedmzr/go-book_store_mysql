package main

import (
	"github.com/gorilla/mux"
	"github.com/saeedmzr/go-simple_book_store/pkg/routes"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	err := http.ListenAndServe(":8080", r) // Updated the listen address
	if err != nil {
		// Handle the error
		panic(err)
	}
}
