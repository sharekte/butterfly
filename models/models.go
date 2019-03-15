package models

import (
	. "spring/types"
)

// DBHandler 数据库操作接口
type (
	DBHandler interface {
		Create() (completed bool, err error)
		FindByID(id int64) (record DBHandler, err error)
		FindByName(name string) (record DBHandler, err error)
		FindAll() (records []DBHandler, err error)
		DeleteByID(id int64) (completed bool, err error)
		DeleteByIDS(ids Int64Values) (completed bool, err error)
		UpdateByID(id int64, data DBHandler) (completed bool, err error)
		BatchUpdateByID(recordValue DynamicValues, id Int64Values) (completed bool, err error)
		SetDefault()
		GetID() int64
		GetActive() int
		getModelName() string
		getPtr() DBHandler
	}
)
