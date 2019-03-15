package service

import (
	"github.com/pkg/errors"
	. "spring/models"
	. "spring/types"
)

type (
	Service struct {
		command Command
	}
)

// Command子服务接口
// 子服务定义示例如下
/*type (
	TempSRV struct {
		model DBHandler ->数据库操作接口，位于models下
		cache Cacher ->缓存操作接口，位于models下
	}
)*/
type (
	Command interface {
		create() (completed bool, err error)
		getByID(id int64) (record DBHandler, err error)
		getByName(name string) (record DBHandler, err error)
		getAll() (records []DBHandler, err error)
		deleteByID(id int64) (completed bool, err error)
		deleteByIDS(ids Int64Values) (completed bool, err error)
		updateByID(id int64) (completed bool, err error)
		batchUpdateByID(recordValue DynamicValues, id Int64Values) (completed bool, err error)
		Init() (srv Command)
		InitWithJson(jsonData []byte) (srv Command, message Message, err error)
		verify(model DBHandler) Message
		Register() (srv Service)
		showData() (data DBHandler)
	}
)

// 服务暴露给外部的接口
func (s Service) Create() (completed bool, err error) {
	if s.command == nil {
		return false, errors.Errorf("调用的服务不存在")
	}
	return s.command.create()
}

func (s Service) GetByID(id int64) (record DBHandler, err error) {
	if s.command == nil {
		return nil, errors.Errorf("调用的服务不存在")
	}
	return s.command.getByID(id)
}

func (s Service) GetByName(name string) (record DBHandler, err error) {
	if s.command == nil {
		return nil, errors.Errorf("调用的服务不存在")
	}
	return s.command.getByName(name)
}

func (s Service) GetAll() (records []DBHandler, err error) {
	if s.command == nil {
		return nil, errors.Errorf("调用的服务不存在")
	}
	return s.command.getAll()
}

func (s Service) DeleteByID(id int64) (completed bool, err error) {
	if s.command == nil {
		return false, errors.Errorf("调用的服务不存在")
	}
	return s.command.deleteByID(id)
}

func (s Service) DeleteByIDS(ids Int64Values) (completed bool, err error) {
	if s.command == nil {
		return false, errors.Errorf("调用的服务不存在")
	}
	return s.command.deleteByIDS(ids)
}

func (s Service) UpdateByID(id int64) (completed bool, err error) {
	if s.command == nil {
		return false, errors.Errorf("调用的服务不存在")
	}
	return s.command.updateByID(id)
}

func (s Service) BatchUpdateByID(recordValue DynamicValues, id Int64Values) (completed bool, err error) {
	if s.command == nil {
		return false, errors.Errorf("调用的服务不存在")
	}
	return s.command.batchUpdateByID(recordValue, id)
}

func (s Service) ShowData() DBHandler {
	return s.command.showData()
}
