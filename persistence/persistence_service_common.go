package persistence

import "reflect"

type PersistenceServiceCommon struct {
	modelType reflect.Type

	findOneLoad  func(id string, value interface{}) error
	findManyLoad func(params QueryParams, values interface{}) error
}

func NewPersistenceServiceCommon(modelType reflect.Type, findOneLoad func(id string, value interface{}) error, findManyLoad func(params QueryParams, values interface{}) error) PersistenceServiceCommon {
	return PersistenceServiceCommon{modelType: modelType, findOneLoad: findOneLoad, findManyLoad: findManyLoad}
}

func (self *PersistenceServiceCommon) GetModelType() reflect.Type {
	return self.modelType
}

func (self *PersistenceServiceCommon) NewModelObj() interface{} {
	return reflect.Zero(self.modelType).Interface()
}

func (self *PersistenceServiceCommon) NewModelObjPtr() interface{} {
	return reflect.New(self.modelType).Interface()
}

func (self *PersistenceServiceCommon) NewModelObjSlice() interface{} {
	var sliceType reflect.Type
	sliceType = reflect.SliceOf(self.modelType)
	return reflect.New(sliceType).Interface()
}

func (self *PersistenceServiceCommon) FindOne(id string) (interface{}, error) {
	ret := self.NewModelObjPtr()
	err := self.findOneLoad(id, ret)
	if err != nil {
		return nil, err
	}
	return reflect.ValueOf(ret).Elem().Interface(), err
}

func (self *PersistenceServiceCommon) FindMany(params QueryParams) (interface{}, error) {
	ret := self.NewModelObjSlice()
	err := self.findManyLoad(params, ret)
	if err != nil {
		return nil, err
	}
	return reflect.ValueOf(ret).Elem().Interface(), err
}
