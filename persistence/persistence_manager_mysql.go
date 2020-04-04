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
	fieldNameToType       map[string]reflect.Type
	fieldNameToNum        map[string]int
	fieldNameToColumnName map[string]string
	idFieldName           string
	idFieldNum            int
}

type PersistenceManagerMySql struct {
	pool      *dbr.Connection
	modelType reflect.Type
	tableName string
	metaData  *mySqlMetaData
}

func (self *PersistenceManagerMySql) FindOne(id string) (interface{}, error) {
	panic("implement me")
}

func (self *PersistenceManagerMySql) FindMany(params QueryParams) (interface{}, error) {
	panic("implement me")
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
	for _, columnName := range self.metaData.fieldNameToColumnName {
		if columnName != self.metaData.idFieldName {
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

func (self *PersistenceManagerMySql) buildMetaData() {
	self.metaData = &mySqlMetaData{
		fieldNameToType:       map[string]reflect.Type{},
		fieldNameToNum:        map[string]int{},
		fieldNameToColumnName: map[string]string{},
	}
	for i := 0; i < self.modelType.NumField(); i++ {
		field := self.modelType.Field(i)
		dbTag := field.Tag.Get("db")
		if dbTag == "" || dbTag == "-" {
			continue
		}
		self.metaData.fieldNameToColumnName[field.Name] = dbTag
		self.metaData.fieldNameToType[field.Name] = field.Type
		self.metaData.fieldNameToNum[field.Name] = i
		if self.metaData.idFieldName == "" && strings.EqualFold(field.Name, "id") || strings.EqualFold(field.Name, "_id") {
			self.metaData.idFieldName = field.Name
			self.metaData.idFieldNum = i
		}
	}
}
