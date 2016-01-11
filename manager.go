package goresource

import (
	"io"
	"net/url"

	"goresource/store"
)

// ResourceManager is an interface implemented by managers for entities.
type ResourceManager interface {
	GetName() string
	New() Entity
	GetEntity(id string, query url.Values) (interface{}, error)
	CreateEntity(entity Entity, query url.Values) (interface{}, error)
	ListEntities(query url.Values) ([]interface{}, error)
	UpdateEntity(id string, entity Entity, query url.Values) (interface{}, error)
	DeleteEntity(id string, query url.Values) error
	ParseJSON(io.ReadCloser) (Entity, error)
}

// DefaultManager is a default implementation for ResourceManager.
// It implements defaults for all methods except New and ParseJSON.
type DefaultManager struct {
	// Name is used a prefix for routes as well as the database collection name.
	Name  string
	Store store.Store
}

// NewDefaultManager initializes and returns a DefaultManager.
func NewDefaultManager(name string, store store.Store) DefaultManager {
	return DefaultManager{name, store}
}

// GetName returns the name for this DefaultManager.
func (manager DefaultManager) GetName() string {
	return manager.Name
}

// GetEntity fetches a single resource entity with the given id.
func (manager DefaultManager) GetEntity(id string, _ url.Values) (interface{}, error) {
	var result interface{}
	if err := manager.Store.GetEntity(manager.Name, id, result); err != nil {
		return nil, err
	}
	return result, nil
}

// CreateEntity persists the given entity.
func (manager DefaultManager) CreateEntity(e Entity, _ url.Values) (interface{}, error) {
	var result interface{}
	if err := manager.Store.CreateEntity(manager.Name, e, result); err != nil {
		return nil, err
	}
	return result, nil
}

// ListEntities fetches all resource entities.
func (manager DefaultManager) ListEntities(_ url.Values) ([]interface{}, error) {
	var result []interface{}
	if err := manager.Store.ListEntities(manager.Name, nil, result); err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateEntity persists changes to the given entity with the given id.
func (manager DefaultManager) UpdateEntity(id string, e Entity, _ url.Values) (interface{}, error) {
	var result interface{}
	if err := manager.Store.UpdateEntity(manager.Name, id, e, result); err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteEntity removes a single entity with the given id.
func (manager DefaultManager) DeleteEntity(id string, _ url.Values) error {
	return manager.Store.DeleteEntity(manager.Name, id)
}
