package persistence

import "reflect"

type QueryParams struct {
	Ands   []interface{} //relational expression array (TBD)
	Ors    []interface{} //relational expression array (TBD)
	Limit  int
	Offset int
}

//validation can be done with something like https://github.com/asaskevich/govalidator or https://github.com/go-ozzo/ozzo-validation
type ValidationErrors struct {
	Errors map[string][]string
}

func (self ValidationErrors) HasErrors() bool {
	return true
}

type PersistenceManager interface {
	FindOne(id string) (interface{}, error)
	FindMany(params QueryParams) (interface{}, error)
	CreateOne(obj interface{})
	//CreateMany(objs interface{})
	DeleteOne(id string) (int, error)
	//DeleteMany(params QueryParams) (int, error)
	UpdateOne(obj interface{}) (int, error)
	//UpdateMany(objs interface{}) (int, error)

	Validate(obj interface{}) (ValidationErrors, error)
	GetEntityType() reflect.Type
}
