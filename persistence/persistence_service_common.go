package persistence

import "reflect"

type PersistenceManagerCommon struct {
	modelType reflect.Type
}

func (self *PersistenceManagerCommon) GetModelType() reflect.Type {
	return self.modelType
}

func (self *PersistenceManagerCommon) NewModelObj() interface{} {
	return reflect.Zero(self.modelType).Interface()
}

func (self *PersistenceManagerCommon) NewModelObjPtr() interface{} {
	return reflect.New(self.modelType).Interface()
}

func (self *PersistenceServiceMySql) NewModelObjSlice() interface{} {
	var sliceType reflect.Type
	sliceType = reflect.SliceOf(self.modelType)
	return reflect.New(sliceType).Interface()
}

func (self *PersistenceServiceMySql) FindOne(id string) (interface{}, error) {
	ret := self.NewModelObjPtr()
	err := self.FindOneLoad(id, ret)
	if err != nil {
		return nil, err
	}
	return reflect.ValueOf(ret).Elem().Interface(), err
}

func (self *PersistenceServiceMySql) FindMany(params QueryParams) (interface{}, error) {
	ret := self.NewModelObjSlice()
	err := self.FindManyLoad(params, ret)
	if err != nil {
		return nil, err
	}
	return reflect.ValueOf(ret).Elem().Interface(), err
}
