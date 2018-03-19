// package store implements a database store for goresource.
package store

import "net/url"

// Store iterface is implemented by database stores.
type Store interface {
	GetEntity(name string, id string, result interface{}) error
	CreateEntity(name string, data interface{}, result interface{}) error
	ListEntities(name string, filters url.Values, result interface{}) error
	UpdateEntity(name string, id string, data interface{}, result interface{}) error
	DeleteEntity(name string, id string) error
	Close()
}
