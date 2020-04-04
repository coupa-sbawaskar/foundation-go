package persistence

import "reflect"

type QueryOperatorType string

const (
	QUERY_OPERATOR_EQ       = "EQ"
	QUERY_OPERATOR_NEQ      = "NEQ"
	QUERY_OPERATOR_CONTAINS = "CONTAINS"
	QUERY_OPERATOR_IN       = "IN"
	QUERY_OPERATOR_GT       = "GT"
	QUERY_OPERATOR_GTE      = "GTE"
	QUERY_OPERATOR_LT       = "LT"
	QUERY_OPERATOR_LTE      = "LTE"
)

type QueryExpression struct {
	ColumnName string
	Operator   QueryOperatorType
	Value      string
}

type QueryParams struct {
	//Ands   []interface{} //relational expression array (TBD)
	//Ors    []interface{} //relational expression array (TBD)
	Operands []QueryExpression
	Limit    uint64
	Offset   uint64
	Order    [][2]string
}

//validation can be done with something like https://github.com/asaskevich/govalidator or https://github.com/go-ozzo/ozzo-validation
type ValidationErrors struct {
	Errors map[string][]string
}

func (self ValidationErrors) HasErrors() bool {
	return true
}

type PersistenceManager interface {
	FindOne(id string, obj interface{}) error
	FindMany(params QueryParams, values interface{}) error
	CreateOne(obj interface{}) error
	//CreateMany(objs interface{})
	DeleteOne(id string) (int, error)
	//DeleteMany(params QueryParams) (int, error)
	UpdateOne(obj interface{}) (int, error)
	//UpdateMany(objs interface{}) (int, error)

	Validate(obj interface{}) (ValidationErrors, error)
	GetEntityType() reflect.Type
}
