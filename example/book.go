package main

import (
	"encoding/json"
	"fmt"
	"io"

	"goresource"
	"goresource/store"
)

type Book struct {
	ID   string `json:"id" bson:"_id"`
	ISBN string `json:"isbn" bson:"isbn"`
	Name string `json:"name" bson:"name"`
}

func (b Book) HasId() bool {
	if b.ID != "" {
		return true
	}
	return false
}

func (b Book) GetId() string {
	return b.ID
}

type BookManager struct {
	goresource.DefaultManager
}

func NewBookManager(name string, store store.Store) *BookManager {
	return &BookManager{goresource.NewDefaultManager(name, store)}
}

func (manager *BookManager) New() goresource.Entity {
	return &Book{}
}

func (manager *BookManager) ParseJSON(data io.ReadCloser) (goresource.Entity, error) {
	book, ok := manager.New().(*Book)
	if !ok {
		return nil, fmt.Errorf("error creating new book.")
	}
	decoder := json.NewDecoder(data)
	err := decoder.Decode(book)
	if err != nil {
		return nil, err
	}
	return book, nil
}
