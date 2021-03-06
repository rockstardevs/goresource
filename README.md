# goresource [![Build Status](https://img.shields.io/travis/rockstardevs/goresource/master.svg?maxAge=3600&style=flat-square)](https://travis-ci.org/rockstardevs/goresource) [![codecov](https://img.shields.io/codecov/c/github/rockstardevs/goresource/master.svg?maxAge=3600&style=flat-square)](https://codecov.io/gh/rockstardevs/goresource) [![coveralls](https://img.shields.io/coveralls/rockstardevs/goresource/master.svg?maxAge=3600&style=flat-square)](https://coveralls.io/github/rockstardevs/goresource?branch=master) 

goresource provides a micro framework for easy implementation of RESTful APIs with golang.

## Overview

goresource comes with a few basic structs and interfaces.

- An **Entity** is an interface for any type that needs to be persisted and operated on.
- A **Resource** is a struct corresponding each entity that can do http request handling.
- A **ResourceManager** ties together a **Resource** and an **Entity**.

### Example

Let's say we have a Book type.

```go
type Book {
  ID   string `json:"id"`
  ISBN string `json:"isbn"`
  Name string `json:"name"`
}
```

To implement a RESTful API for the Book type, first we satisfy the Entity interface for Book.

```go
func (b Book) HasId() bool {
  if b.ID != "" {
      return true
    }
    return false
}

func (b Book) GetId() string {
  return b.ID
}
```

Next we write a ResourceManager for Book.

```go
type BookManager struct {
  goresource.DefaultManager
}

func NewBookManager(name string, store store.Store) *BookManager {
  return &BookManager{goresource.NewDefaultManager(name, store)}
}
```

BookManager should satify the ResourceManager interface

```go
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
```

Finally we create a new Resource with the BookManager

```go
router := mux.NewRouter().PathPrefix("/api").Subrouter()
store, _ := store.NewMongoStore("localhost:27017", "booksdb")
defer store.Close()

manager := NewBookManager("books", store)
bookResource := goresource.NewResource(manager, router)
```

The router can them be used to serve the REST API endpoints.

```go
http.Handle("/", router)
http.ListenAndServe(":8080", nil)
```

See the example directory for the full example code.

## Installation

```sh
go get github.com/rockstardevs/goresource
```

## Test

tests are written using [ginkgo](http://github.com/onsi/ginkgo).

```sh
go get github.com/onsi/ginkgo/ginkgo github.com/onsi/gomega
```

To run all tests recursively

```sh
ginkgo -r
```
## Updating Mocks

```sh
mockgen -package mocks -destination mocks/store.go goresource/store Store
mockgen -package mocks -destination mocks/manager.go goresource ResourceManager
```

## TODO/What could be better

- Implement PATCH and OPTIONS
- Remove hard dependency on mux.Router and accept any Router.
- Implement addition stores, currently only MongoDB is implemented.

## Contributing

Pull requests are welcome. Please ensure to include tests.
