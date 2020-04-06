package rest

import (
	"fmt"
	"github.com/coupa/foundation-go/persistence"
	"reflect"
	"strings"
)

type PersistenceServiceRest struct {
	persistence.PersistenceServiceCommon

	baseUrl    string
	RestClient *RestClient
}

var _ persistence.PersistenceService = (*PersistenceServiceRest)(nil)

func NewPersistenceServiceRest(baseUrl string, modelType reflect.Type) (*PersistenceServiceRest, error) {
	ret := &PersistenceServiceRest{
		baseUrl: baseUrl,
	}
	ret.PersistenceServiceCommon = persistence.NewPersistenceServiceCommon(modelType, ret.FindOneLoad, ret.FindManyLoad)
	return ret, nil
}

func (self *PersistenceServiceRest) FindOneLoad(id string, value interface{}) error {
	resp := self.restClient().GetObject(self.baseUrl+"/"+id, value)
	return resp.Error
}

func (self *PersistenceServiceRest) FindManyLoad(params persistence.QueryParams, values interface{}) error {
	url := strings.Builder{}
	url.WriteString(self.baseUrl)
	url.WriteString("?")
	if params.Limit > 0 {
		url.WriteString(fmt.Sprintf("&limit=%v", params.Limit))
	}
	if params.Offset > 0 {
		url.WriteString(fmt.Sprintf("&offset=%v", params.Offset))
	}
	for _, query := range params.Operands {
		url.WriteString(fmt.Sprintf("q[%v_%v]=%v", query.Key, query.Operator, query.Value))
	}
	for _, order := range params.Order {
		url.WriteString(fmt.Sprintf("&order=%v,%v", order.Key, order.Direction))
	}
	resp := self.restClient().GetObject(url.String(), values)
	return resp.Error
}

func (self *PersistenceServiceRest) CreateOne(obj interface{}) error {
	resp := self.restClient().PostObject(self.baseUrl, obj)
	return resp.Error
}

func (self *PersistenceServiceRest) DeleteOne(id string) (bool, error) {
	resp := self.restClient().DeleteObject(self.baseUrl + "/" + id)
	return resp.Error == nil, resp.Error
}

func (self *PersistenceServiceRest) UpdateOne(id string, obj interface{}) (bool, error) {
	resp := self.restClient().PutObject(self.baseUrl+"/"+id, obj)
	return resp.Error == nil, resp.Error
}

func (self *PersistenceServiceRest) Validate(obj interface{}) (persistence.ValidationErrors, error) {
	panic("implement me")
}

func (self *PersistenceServiceRest) restClient() *RestClient {
	if self.RestClient == nil {
		return &RestClient{}
	}
	return self.RestClient
}
