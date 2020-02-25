package persistence

import (
	"reflect"
)

var _ PersistenceManager = (*PersistenceManagerMongo)(nil)

type PersistenceManagerMongo struct {
	Session interface{}
}

func (self PersistenceManagerMongo) FindOne(id string) (interface{}, error) {
	panic("implement me")
}

func (self PersistenceManagerMongo) FindMany(params QueryParams) (interface{}, error) {
	panic("implement me")
}

func (self PersistenceManagerMongo) CreateOne(obj interface{}) {
	panic("implement me")
}

func (self PersistenceManagerMongo) DeleteOne(id string) (int, error) {
	panic("implement me")
}

func (self PersistenceManagerMongo) DeleteMany(params QueryParams) (int, error) {
	panic("implement me")
}

func (self PersistenceManagerMongo) UpdateOne(id string) (int, error) {
	panic("implement me")
}

func (self PersistenceManagerMongo) UpdateMany(params QueryParams) (int, error) {
	panic("implement me")
}

func (self PersistenceManagerMongo) GetEntityType() reflect.Type {
	panic("implement me")
}
