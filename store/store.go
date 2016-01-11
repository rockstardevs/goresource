// package store implements a database store for goresource.
package store

// Store iterface is implemented by database stores.
type Store interface {
	GetEntity(name string, id string) (interface{}, error)
	CreateEntity(name string, data interface{}) (interface{}, error)
	ListEntities(name string, filters interface{}) ([]interface{}, error)
	UpdateEntity(name string, id string, data interface{}) (interface{}, error)
	DeleteEntity(name string, id string) error
	Close()
}
