package persistence

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gocraft/dbr/v2"
	"reflect"
	"strconv"
	"strings"
)

type PersistenceServiceMySql struct {
	PersistenceServiceCommon

	connection *dbr.Connection
	tableName  string
	metaData   *mySqlMetaData
}

var _ PersistenceService = (*PersistenceServiceMySql)(nil)

func NewPersistenceServiceMySql(dsn string, tableName string, modelType reflect.Type) (*PersistenceServiceMySql, error) {
	connection, err := dbr.Open("mysql", dsn, nil)
	if err != nil {
		return nil, err
	}
	ret := &PersistenceServiceMySql{
		connection: connection,
		tableName:  tableName,
	}
	ret.PersistenceServiceCommon = NewPersistenceServiceCommon(modelType, ret.FindOneLoad, ret.FindManyLoad)
	ret.buildMetaData()
	return ret, nil
}

type mySqlMetaData struct {
	columnNameToType map[string]reflect.Type
	columnNameToNum  map[string]int
	//columnNameToFieldName map[string]string
	idColumnName string
	idFieldNum   int
}

func (self *PersistenceServiceMySql) FindOneLoad(id string, value interface{}) error {
	session := self.connection.NewSession(nil)
	query := session.Select("*").From(self.tableName)
	self.queryId(query, id)
	_, err := query.Load(value)
	return err
}

func (self *PersistenceServiceMySql) FindManyLoad(params QueryParams, values interface{}) error {
	session := self.connection.NewSession(nil)
	query := session.Select("*").From(self.tableName)
	if params.Limit > 0 {
		query.Limit(params.Limit)
	}
	if params.Offset > 0 {
		query.Offset(params.Offset)
	}
	for _, op := range params.Operands {
		var operator string
		var value interface{} = op.Value
		switch op.Operator {
		case QUERY_OPERATOR_EQ:
			operator = "="
			value, _ = strconv.ParseBool(op.Value)
		case QUERY_OPERATOR_NEQ:
			operator = "!="
			value, _ = strconv.ParseBool(op.Value)
		case QUERY_OPERATOR_GT:
			operator = ">"
		case QUERY_OPERATOR_GTE:
			operator = ">="
		case QUERY_OPERATOR_LT:
			operator = "<"
		case QUERY_OPERATOR_LTE:
			operator = "<="
		case QUERY_OPERATOR_CONTAINS:
			operator = "like"
			value = "%" + op.Value + "%"
		case QUERY_OPERATOR_IN:
			operator = "in"
			value = strings.Split(op.Value, ",")
		case QUERY_OPERATOR_STARTS_WITH:
			operator = "like"
			value = op.Value + "%"
		case QUERY_OPERATOR_ENDS_WITH:
			operator = "like"
			value = "%" + op.Value
		default:
			return fmt.Errorf("unknown operator '%v'", op.Operator)
		}

		query.Where(fmt.Sprintf("%s %s ?", op.Key, operator), value)
	}
	for _, order := range params.Order {
		if order.Direction == ORDER_DIRECTION_ASC {
			query.OrderAsc(order.Key)
		} else {
			query.OrderDesc(order.Key)
		}
	}
	_, err := query.Load(values)
	return err
}

func (self *PersistenceServiceMySql) CreateOne(obj interface{}) error {
	objType := reflect.TypeOf(obj)
	if objType.Kind() != reflect.Ptr {
		return fmt.Errorf("obj must be a pointer")
	}

	session := self.connection.NewSession(nil)
	query := session.InsertInto(self.tableName)
	columns := []string{}
	for columnName, _ := range self.metaData.columnNameToNum {
		if columnName != self.metaData.idColumnName {
			columns = append(columns, columnName)
		}
	}
	_, err := query.Columns(columns...).Record(obj).Exec()
	return err
}

func (self *PersistenceServiceMySql) DeleteOne(id string) (bool, error) {
	session := self.connection.NewSession(nil)
	res, err := session.DeleteFrom(self.tableName).Where(fmt.Sprintf("%s=?", self.metaData.idColumnName), id).Exec()
	if err != nil {
		return false, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return false, err
	}
	return rowsAffected == 1, nil
}

func (self *PersistenceServiceMySql) UpdateOne(id string, obj interface{}) (bool, error) {
	objType := reflect.TypeOf(obj)
	if objType.Kind() != reflect.Ptr {
		return false, fmt.Errorf("obj must be a pointer")
	}

	session := self.connection.NewSession(nil)
	query := session.Update(self.tableName)
	objValue := reflect.ValueOf(obj)
	for columnName, columnNum := range self.metaData.columnNameToNum {
		if columnName != self.metaData.idColumnName {
			query.Set(columnName, objValue.Elem().Field(columnNum).Interface())
		}
	}
	res, err := query.Where(fmt.Sprintf("%s=?", self.metaData.idColumnName), id).Exec()
	if err != nil {
		return false, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return false, err
	}
	return rowsAffected == 1, nil
}

func (self *PersistenceServiceMySql) Validate(obj interface{}) (ValidationErrors, error) {
	return Validate(obj)
}

func (self *PersistenceServiceMySql) queryId(query *dbr.SelectStmt, id string) {
	query.Where(fmt.Sprintf("%s=?", self.metaData.idColumnName), id)
}

func (self *PersistenceServiceMySql) buildMetaData() {
	self.metaData = &mySqlMetaData{
		columnNameToType: map[string]reflect.Type{},
		columnNameToNum:  map[string]int{},
		//fieldNameToColumnName: map[string]string{},
	}
	for i := 0; i < self.modelType.NumField(); i++ {
		field := self.modelType.Field(i)
		dbTag := field.Tag.Get("db")
		if dbTag == "" || dbTag == "-" {
			continue
		}
		//self.metaData.fieldNameToColumnName[field.Name] = dbTag
		self.metaData.columnNameToType[dbTag] = field.Type
		self.metaData.columnNameToNum[dbTag] = i
		if self.metaData.idColumnName == "" && (strings.EqualFold(field.Name, "id") || strings.EqualFold(field.Name, "_id")) {
			self.metaData.idColumnName = dbTag
			self.metaData.idFieldNum = i
		}
	}
}

func (self *PersistenceServiceMySql) GetConnection() *dbr.Connection {
	return self.connection
}
