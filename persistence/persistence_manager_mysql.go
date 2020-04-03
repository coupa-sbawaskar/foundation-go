package persistence

import (
	"reflect"
)

var _ PersistenceManager = (*PersistenceManagerMySql)(nil)

type PersistenceManagerMySql struct {
	session    interface{}
	entityType reflect.Type
}

func (self PersistenceManagerMySql) FindOne(id string) (interface{}, error) {
	panic("implement me")
}

func (self PersistenceManagerMySql) FindMany(params QueryParams) (interface{}, error) {
	panic("implement me")
}

func (self PersistenceManagerMySql) CreateOne(obj interface{}) {
	panic("implement me")
}

func (self PersistenceManagerMySql) DeleteOne(id string) (int, error) {
	panic("implement me")
}

func (self PersistenceManagerMySql) UpdateOne(obj interface{}) (int, error) {
	panic("implement me")
}

func (self PersistenceManagerMySql) Validate(obj interface{}) (ValidationErrors, error) {
	panic("implement me")
}

func (self PersistenceManagerMySql) GetEntityType() reflect.Type {
	panic("implement me")
}

func (self PersistenceManagerMySql) getSession() {

}
