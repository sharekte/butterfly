package db

import (
	"fmt"
	"strings"
)

// 数据库查询参数自定义类型
type (
	Int64Values   []int64
	StrValues     []string
	IntValues     []int
	DynamicValues []interface{}
	Record        map[string]interface{}
	Message       map[string]interface{}
)

type (
	Argument interface {
		UnPack() (values []interface{})
		Pack2Slice() (valueSet [][]interface{})
		Get(item interface{}) (exist bool, index int, value interface{})
		Len() int
	}
)

func (i Int64Values) UnPack() (values []interface{}) {
	for _, v := range i {
		values = append(values, v)
	}
	return
}

func (i Int64Values) Pack2Slice() (valueSet [][]interface{}) {
	valueSet = append(valueSet, i.UnPack())
	return
}

func (i Int64Values) Get(item interface{}) (exist bool, index int, value interface{}) {
	item4Int64, ok := item.(int64)
	if ok == false {
		return
	}
	for j, v := range i {
		if v == item4Int64 {
			exist, index, value = true, j, item4Int64
			break
		}
	}
	return

}

func (i Int64Values) Len() int {
	return len(i)
}

func (s StrValues) UnPack() (values []interface{}) {
	for _, v := range s {
		values = append(values, v)
	}
	return
}

func (s StrValues) Pack2Slice() (valueSet [][]interface{}) {
	valueSet = append(valueSet, s.UnPack())
	return
}

func (s StrValues) Get(item interface{}) (exist bool, index int, value interface{}) {
	item4Str, ok := item.(string)
	if ok == false {
		return
	}
	for i, v := range s {
		if v == item4Str {
			exist, index, value = true, i, item4Str
			break
		}
	}
	return
}

func (s StrValues) Len() int {
	return len(s)
}

func (i IntValues) UnPack() (values []interface{}) {
	for _, v := range i {
		values = append(values, v)
	}
	return
}

func (i IntValues) Pack2Slice() (valueSet [][]interface{}) {
	valueSet = append(valueSet, i.UnPack())
	return
}

func (i IntValues) Get(item interface{}) (exist bool, index int, value interface{}) {
	item4Int, ok := item.(int)
	if ok == false {
		return
	}
	for i, v := range i {
		if v == item4Int {
			exist, index, value = true, i, item4Int
			break
		}
	}
	return

}

func (i IntValues) Len() int {
	return len(i)
}

func (d DynamicValues) UnPack() (values []interface{}) {
	return d
}

func (d DynamicValues) Pack2Slice() (valueSet [][]interface{}) {
	valueSet = append(valueSet, d)
	return
}

func (d DynamicValues) Get(item interface{}) (exist bool, index int, value interface{}) {
	for i, v := range d {
		if v == item {
			exist, index, value = true, i, item
			break
		}
	}
	return

}

func (d DynamicValues) Len() int {
	return len(d)
}

// 只返回map的值
func (r Record) UnPack() (values []interface{}) {
	for _, v := range r {
		values = append(values, v)
	}
	return
}

// 返回map键，值的二维切片
func (r Record) Pack2Slice() (valueSet [][]interface{}) {
	keys := make([]interface{}, 0)
	values := make([]interface{}, 0)
	for k, v := range r {
		keys = append(keys, k)
		values = append(values, v)
	}
	valueSet = append(valueSet, keys, values)
	return

}

func (r Record) Get(item interface{}) (exist bool, index int, value interface{}) {
	item4Str, ok := item.(string)
	if ok == false {
		return
	}
	for k, v := range r {
		if k == item4Str {
			exist, index, value = true, 0, v
			break
		}
	}
	return
}

func (r Record) Len() int {
	return len(r)
}

func (r Record) Pack4SQL() (kvSet string, fieldSet string, valueSet string, values DynamicValues) {
	var num int = 0
	var tempChar = "%s%d"
	var tempKVSet []string
	var fields []string
	var tempValues []string
	for k, v := range r {
		tempChar = fmt.Sprintf(tempChar, "%s", num)
		subKV := fmt.Sprintf("%s=%s", k, tempChar)
		tempKVSet = append(tempKVSet, subKV)
		values = append(values, v)
		fields = append(fields, k)
		tempValues = append(tempValues, tempChar)
		num++
	}
	kvSet = strings.Join(tempKVSet, ",")
	fieldSet = strings.Join(fields, ",")
	valueSet = strings.Join(tempValues, ",")
	return
}
