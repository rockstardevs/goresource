package store

import (
	"github.com/golang/glog"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// MongoStore is a store implementation using mongodb as the database.
type MongoStore struct {
	session *mgo.Session
	db      *mgo.Database
}

// NewMongoStore returns an initialized store or any initialization errors.
func NewMongoStore(addr string, database string) (Store, error) {
	session, err := mgo.Dial(addr)
	if err != nil {
		return nil, err
	}
	return &MongoStore{
		session: session,
		db:      session.DB(database),
	}, nil
}

// ListEntities queries and returns all entities matching the given filters.
func (s *MongoStore) ListEntities(name string, filters interface{}) ([]interface{}, error) {
	entities := make([]interface{}, 0)
	err := s.db.C(name).Find(filters).All(&entities)
	if err != nil {
		glog.Errorf("error listing entities of %s - %s", name, err)
		return nil, err
	}
	return entities, nil
}

// ListEntities fetches a specific entity with the given id.
func (s *MongoStore) GetEntity(name string, id string) (interface{}, error) {
	var entity interface{}
	entityId := bson.ObjectIdHex(id)
	err := s.db.C(name).FindId(entityId).One(&entity)
	if err != nil {
		return nil, err
	}
	return entity, nil
}

// CreateEntity persists a new entity with the given data.
func (s *MongoStore) CreateEntity(name string, data interface{}) (interface{}, error) {
	var result interface{}
	err := s.db.C(name).Insert(data)
	if err != nil {
		return nil, err
	}
	err = s.db.C(name).Find(data).One(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateEntity updates a specific entity corresponding the given id, with the given data.
func (s *MongoStore) UpdateEntity(name string, id string, data interface{}) (interface{}, error) {
	entityId := bson.ObjectIdHex(id)
	if err := s.db.C(name).UpdateId(entityId, data); err != nil {
		return nil, err
	}
	return data, nil
}

// DeleteEntity removes a specific entity with the given id.
func (s *MongoStore) DeleteEntity(name string, id string) error {
	return s.db.C(name).RemoveId(bson.ObjectIdHex(id))
}

// Close tears down the database connection and closes the session.
func (s *MongoStore) Close() {
	s.session.Close()
}
