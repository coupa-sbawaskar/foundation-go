package persistence

import "reflect"

type PersistenceService interface {
	FindOne(id string) (interface{}, error)
	FindOneLoad(id string, value interface{}) error
	FindMany(params QueryParams) (values interface{}, err error)
	FindManyLoad(params QueryParams, values interface{}) error
	CreateOne(obj interface{}) error
	DeleteOne(id string) (bool, error)
	UpdateOne(id string, obj interface{}) (bool, error)
	NewModelObj() interface{}
	NewModelObjPtr() interface{}

	Validate(obj interface{}) (ValidationErrors, error)
	GetModelType() reflect.Type

	//todo CreateMany(objs interface{})
	//todo DeleteMany(params QueryParams) (int, error)
	//todo UpdateMany(objs interface{}) (int, error)
}

//validation can be done with something like https://github.com/asaskevich/govalidator or https://github.com/go-ozzo/ozzo-validation
type ValidationErrors struct {
	Errors map[string][]string `json:"errors"`
}

func (self *ValidationErrors) HasErrors() bool {
	return self.Errors != nil && len(self.Errors) > 0
}
