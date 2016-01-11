// Package goresource provides a micro REST framework.
package goresource

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rockstardevs/goresource/util"
)

// Entity is an interface implemented by entities we store in the database.
type Entity interface {
	HasId() bool
	GetId() string
}

// Resource provides an abstraction to be able to store and retrieve
// entities via a RESTful api. It is responsible for request handling and
// writing responses. The acutal persistence and any entity specific operations
// are delegated to the corresponding ResourceManager. This decouples request
// handing and persistence from specific entity types.
type Resource struct {
	manager ResourceManager
}

// NewResource instantiates a Resource and binds routes to the given mux router,
// to serve the api end points specific to this resource.
func NewResource(m ResourceManager, router *mux.Router) *Resource {
	r := &Resource{manager: m}
	router.Handle(fmt.Sprintf("/%s", m.GetName()), r)
	router.Handle(fmt.Sprintf("/%s/{id}", m.GetName()), r)
	return r
}

// ServeHTTP is the main http handler that handles all api request for this resource.
// It delegates based on HTTP Method to other methods of this resource.
func (r Resource) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		r.Get(rw, req)
	case "POST":
		r.PostOrPut(rw, req)
	case "PUT":
		r.PostOrPut(rw, req)
	case "DELETE":
		r.Delete(rw, req)
	case "HEAD":
		r.Head(rw, req)
	case "PATCH":
		r.Patch(rw, req)
	default:
		r.UnsupportedMethod(rw, req)
	}
}

// get is the common code between get and head requests.
func (r Resource) get(rw http.ResponseWriter, req *http.Request) interface{} {
	var (
		query = req.URL.Query()
		vars  = mux.Vars(req)
		resp  interface{}
		err   error
	)
	id := vars["id"]
	if id != "" {
		resp, err = r.manager.GetEntity(id, query)
	} else {
		resp, err = r.manager.ListEntities(query)
	}
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return nil
	}
	return resp
}

// Get is the delegate http handler for get requests for this resource.
func (r Resource) Get(rw http.ResponseWriter, req *http.Request) {
	resp := r.get(rw, req)
	if resp != nil {
		util.WriteJSON(resp, rw)
	}
}

// Head is the delegate http handler for head requests for this resource.
func (r Resource) Head(rw http.ResponseWriter, req *http.Request) {
	if r.get(rw, req) != nil {
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
	}
}

// PostOrPut is the common code between put and post requests.
// TODO: This is partially incorrect according to the spec.
// PUT without an id in the uri is invalid and should return a BadRequest error.
// Fix this to comply to with the spec.
func (r Resource) PostOrPut(rw http.ResponseWriter, req *http.Request) {
	var (
		vars   = mux.Vars(req)
		query  = req.URL.Query()
		entity Entity
		resp   interface{}
		err    error
	)
	if entity, err = r.manager.ParseJSON(req.Body); err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	id := vars["id"]
	if id == "" && entity.HasId() {
		id = entity.GetId()
	}
	if id != "" {
		resp, err = r.manager.UpdateEntity(id, entity, query)
	} else {
		resp, err = r.manager.CreateEntity(entity, query)
	}
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	util.WriteJSON(resp, rw)
}

// Delete is the delegate http handler for delete requests for this resource.
func (r Resource) Delete(rw http.ResponseWriter, req *http.Request) {
	var (
		query = req.URL.Query()
		vars  = mux.Vars(req)
		err   error
	)
	id := vars["id"]
	if id == "" {
		http.Error(rw, "Invalid Id", http.StatusBadRequest)
		return
	}
	if err = r.manager.DeleteEntity(id, query); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	rw.WriteHeader(http.StatusNoContent) // Status 204 OK
}

// Patch is the delegate http handler for patch requests for this resource.
func (r Resource) Patch(rw http.ResponseWriter, req *http.Request) {
	// TODO: Implement this.
	http.Error(rw, "Method Not Supported", http.StatusNotImplemented)
}

// UnsupportedMethod is the delegate http handler for unknown requests types.
// This is the catch-all when request method is not known.
func (r Resource) UnsupportedMethod(rw http.ResponseWriter, req *http.Request) {
	http.Error(rw, "Method Not Supported", http.StatusNotImplemented)
}
