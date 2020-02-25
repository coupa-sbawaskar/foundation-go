package persistence

import (
	"reflect"
)

var _ PersistenceManager = (*PersistenceManagerMySql)(nil)

type PersistenceManagerMySql struct {
	session    interface{}
	entityType reflect.Type
}

func NewPersistenceManagerMySql(session interface{}, entityType reflect.Type) PersistenceManagerMySql {
	return PersistenceManagerMySql{session: session, entityType: entityType}
}

func (self PersistenceManagerMySql) CreateOne(obj interface{}) {
	panic("implement me")
}

func (self PersistenceManagerMySql) GetEntityType() reflect.Type {
	panic("implement me")
}

func (self PersistenceManagerMySql) FindOne(id string) (interface{}, error) {
	panic("implement me")
}

func (self PersistenceManagerMySql) FindMany(params QueryParams) (interface{}, error) {
	panic("implement me")
}

func (self PersistenceManagerMySql) DeleteOne(id string) (int, error) {
	panic("implement me")
}

func (self PersistenceManagerMySql) DeleteMany(params QueryParams) (int, error) {
	panic("implement me")
}

func (self PersistenceManagerMySql) UpdateOne(id string) (int, error) {
	panic("implement me")
}

func (self PersistenceManagerMySql) UpdateMany(params QueryParams) (int, error) {
	panic("implement me")
}
