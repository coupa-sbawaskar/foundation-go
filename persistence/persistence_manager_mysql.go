package persistence

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gocraft/dbr/v2"
	"reflect"
	"strings"
)

var _ PersistenceManager = (*PersistenceManagerMySql)(nil)

func NewPersistenceManagerMySql(dsn string, tableName string, modelType reflect.Type) (*PersistenceManagerMySql, error) {
	pool, err := dbr.Open("mysql", dsn, nil)
	if err != nil {
		return nil, err
	}
	ret := &PersistenceManagerMySql{
		pool:      pool,
		tableName: tableName,
		modelType: modelType,
	}
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

type PersistenceManagerMySql struct {
	pool      *dbr.Connection
	modelType reflect.Type
	tableName string
	metaData  *mySqlMetaData
}

func (self *PersistenceManagerMySql) FindOne(id string, value interface{}) error {
	session := self.pool.NewSession(nil)
	defer session.Close()
	query := session.Select("*").From(self.tableName)
	self.queryId(query, id)
	_, err := query.Load(value)
	return err
}

func (self *PersistenceManagerMySql) FindMany(params QueryParams, values interface{}) error {
	session := self.pool.NewSession(nil)
	defer session.Close()
	query := session.Select("*").From(self.tableName)
	if params.Limit > 0 {
		query.Limit(params.Limit)
	}
	if params.Offset > 0 {
		query.Offset(params.Offset)
	}
	for _, op := range params.Operands {
		var operator string
		var operand interface{} = op.Value
		switch op.Operator {
		case QUERY_OPERATOR_EQ:
			operator = "="
		case QUERY_OPERATOR_NEQ:
			operator = "!="
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
			operand = "%" + op.Value + "%"
		case QUERY_OPERATOR_IN:
			operator = "in"
			operand = strings.Split(op.Value, ",")
		}
		for _, order := range params.Order {
			if strings.EqualFold("desc", order[1]) {
				query.OrderDesc(order[0])
			} else if strings.EqualFold("asc", order[1]) {
				query.OrderAsc(order[0])
			}
		}
		query.Where(fmt.Sprintf("%s %s ?", op.ColumnName, operator), operand)
		fmt.Print(operator)
		fmt.Print(operand)
	}
	_, err := query.Load(values)
	return err
}

func (self *PersistenceManagerMySql) CreateOne(obj interface{}) error {
	objType := reflect.TypeOf(obj)
	if objType.Kind() != reflect.Ptr {
		return fmt.Errorf("obj must be a pointer")
	}

	session := self.pool.NewSession(nil)
	defer session.Close()
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

func (self *PersistenceManagerMySql) DeleteOne(id string) (int, error) {
	panic("implement me")
}

func (self *PersistenceManagerMySql) UpdateOne(obj interface{}) (int, error) {
	panic("implement me")
}

func (self *PersistenceManagerMySql) Validate(obj interface{}) (ValidationErrors, error) {
	panic("implement me")
}

func (self *PersistenceManagerMySql) GetEntityType() reflect.Type {
	panic("implement me")
}

func (self *PersistenceManagerMySql) queryId(query *dbr.SelectStmt, id string) {
	query.Where(fmt.Sprintf("%s=?", self.metaData.idColumnName), id)
}

func (self *PersistenceManagerMySql) buildMetaData() {
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
