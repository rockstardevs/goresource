package store

import (
	"fmt"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// MongoStore is a store implementation using mongodb as the database.
type MongoStore struct {
	session *mgo.Session
	db      *mgo.Database
}

// NewMongoStore returns an initialized store or any initialization errors.
func NewMongoStore(addr string, database string, timeout time.Duration) (Store, error) {
	var (
		session *mgo.Session
		err     error
	)
	if timeout != 0 {
		info := &mgo.DialInfo{Addrs: []string{addr}, Timeout: timeout, Database: database}
		session, err = mgo.DialWithInfo(info)
	} else {
		session, err = mgo.Dial(addr)
	}
	if err != nil {
		return nil, err
	}
	return &MongoStore{
		session: session,
		db:      session.DB(database),
	}, nil
}

// ListEntities queries and returns all entities matching the given filters.
func (s *MongoStore) ListEntities(name string, filters interface{}, result interface{}) error {
	err := s.db.C(name).Find(filters).All(result)
	if err != nil {
		return err
	}
	return nil
}

// ListEntities fetches a specific entity with the given id.
func (s *MongoStore) GetEntity(name string, id string, result interface{}) error {
	if !bson.IsObjectIdHex(id) {
		return fmt.Errorf("invalid object id %s.", id)
	}
	entityId := bson.ObjectIdHex(id)
	err := s.db.C(name).FindId(entityId).One(result)
	if err != nil {
		return err
	}
	return nil
}

// CreateEntity persists a new entity with the given data.
func (s *MongoStore) CreateEntity(name string, data interface{}, result interface{}) error {
	err := s.db.C(name).Insert(data)
	if err != nil {
		return err
	}
	if err = s.db.C(name).Find(data).One(result); err != nil {
		return err
	}
	return nil
}

// UpdateEntity updates a specific entity corresponding the given id, with the given data.
func (s *MongoStore) UpdateEntity(name string, id string, data interface{}, result interface{}) error {
	entityId := bson.ObjectIdHex(id)
	err := s.db.C(name).UpdateId(entityId, data)
	if err != nil {
		return err
	}
	if err = s.db.C(name).FindId(entityId).One(result); err != nil {
		return err
	}
	return nil
}

// DeleteEntity removes a specific entity with the given id.
func (s *MongoStore) DeleteEntity(name string, id string) error {
	return s.db.C(name).RemoveId(bson.ObjectIdHex(id))
}

// Close tears down the database connection and closes the session.
func (s *MongoStore) Close() {
	if s.session != nil {
		s.session.Close()
	}
}
