# goresource

goresource provides a micro framework for easy implementation of RESTful APIs with golang.

## Overview

goresource comes with a few basic structs and interfaces.

- An **Entity** is an interface for any type that needs to be persisted and operated on.
- A **Resource** is a struct corresponding each entity that can do http request handling.
- A **ResourceManager** ties together a **Resource** and an **Entity**.

### Example

Let's say we have a Book type.

```golang

type Book {
  ID   string `json:"id"`
    ISBN string `json:"isbn"`
  Name string `json:"name"`
}
```

To implement a RESTful API for the Book type, first we satisfy the Entity interface for Book.

```golang

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

```golang

type BookManager struct {
  goresource.DefaultManager
}

func NewBookManager(name string, store store.Store) *BookManager {
  return &BookManager{goresource.NewDefaultManager(name, store)}
}
```

BookManager should satify the ResourceManager interface

```golang

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

```golang

router := mux.NewRouter().PathPrefix("/api").Subrouter()
store := store.NewMongoStore("localhost:27017", "booksdb")
defer store.Close()

manager := NewBookManager("books", store)
bookResource := goresource.NewResource(manager, router)
```

The router can them be used to serve the REST API endpoints.

```golang

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
go get github.com/onsi/ginkgo github.com/onsi/gomega
```

To run all tests recursively

```sh
ginkgo -r
```

## TODO/What could be better

- Implement PATCH and OPTIONS
- Remove hard dependency on mux.Router and accept any Router.
- Implement addition stores, currently only MongoDB is implemented.
- Unit tests for store.
- Setup CI

## Contributing

Pull requests are welcome. Please ensure to include tests.