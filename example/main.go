package main

import (
	"log"
	"net/http"
	"time"

	"goresource"
	"goresource/store"

	"github.com/gorilla/mux"
)

func main() {
	store, err := store.NewMongoStore("localhost:27017", "booksdb", 5*time.Second)
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
