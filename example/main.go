package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rockstardevs/goresource"
	"github.com/rockstardevs/goresource/store"
)

func main() {
	store, err := store.NewMongoStore("localhost:27017", "booksdb")
	if err != nil {
		log.Fatal("error connecting store - %s", err)
	}
	defer store.Close()

	router := mux.NewRouter()
	apirouter := router.PathPrefix("/api").Subrouter()
	manager := NewBookManager("books", store)
	goresource.NewResource(manager, apirouter)

	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}
