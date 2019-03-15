package models

import "time"

// Cacher 针对Redis的缓存操作接口
type (
	Cacher interface {
		HMSet(key string, fields interface{}) (string, error)
		getKeys(keyPattern string) ([]string, error)
		GetRecordByKey(key string) (keys []string, records []interface{}, err error)
		Insert(key string, data interface{}, expired time.Duration) (string, error)
		MultiInsertByMap(data map[string]interface{}) (string, error)
		MultiInsertByStruct(data interface{}) (string, error)
		DelRecord(key string) (bool, error)
		del(key ...string) (bool, error)
		UpdateByKey(src []byte, key string) (bool, error)
		MultiUpdateByKey(keyPattern string, record []byte) (bool, error)
		NewCacheSRV() Cacher
	}
)
