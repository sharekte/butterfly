package models

import (
	. "butterfly/types"
)

// DBHandler 数据库操作接口
type (
	DBHandler interface{
		create() (completed bool, err error)
		findByID(id int64) (record DBHandler, err error)
		findByName(name string) (record DBHandler, err error)
		findAll() (records []DBHandler, err error)
		deleteByID(id int64) (completed bool, err error)
		deleteByIDS(ids Int64Values) (completed bool, err error)
		updateByID(id int64) (completed bool, err error)
		batchUpdateByID(records []DBHandler) (completed bool, err error)
		setDefault()
		getID() int64
		getActive() int
		getModelName() string
		getPtr() DBHandler
		Register()*Model
	}
)

type(
	Model struct {
		action DBHandler
	}
)

func(m *Model)Create() (completed bool, err error){
	return m.action.create()
}

func(m *Model)FindByID(id int64) (record DBHandler, err error){
	return m.action.findByID(id)

}
func(m *Model)FindByName(name string) (record DBHandler, err error){
	return m.action.findByName(name)

}
func(m *Model)FindAll() (records []DBHandler, err error){
	return m.action.findAll()

}
func(m *Model)DeleteByID(id int64) (completed bool, err error){
	return m.action.deleteByID(id)

}
func(m *Model)DeleteByIDS(ids Int64Values) (completed bool, err error){
	return m.action.deleteByIDS(ids)

}
func(m *Model)UpdateByID(id int64) (completed bool, err error){
	return m.action.updateByID(id)

}
func(m *Model)BatchUpdateByID(records []DBHandler) (completed bool, err error){
	return m.action.batchUpdateByID(records)

}
func(m *Model)SetDefault(){
	m.action.setDefault()

}
func(m *Model)GetID() int64{
	return m.action.getID()

}
func(m *Model)GetActive() int{
	m.action.getActive()

}
func(m *Model)getModelName() string{
	return m.action.getModelName()

}
func(m *Model)getPtr() DBHandler{
	return m.action.getPtr()
}

func(m *Model) registerModel(model DBHandler)*Model{
	m.action=model
	return m
}
