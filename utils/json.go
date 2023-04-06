package utils

/**
Copyright (c) [2023] [jaden Hsiao]
[Poscom] is licensed under Mulan PSL v2.
You can use this software according to the terms and conditions of the Mulan PSL v2.
You may obtain a copy of Mulan PSL v2 at:
http://license.coscl.org.cn/MulanPSL2
THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
See the Mulan PSL v2 for more details.
*/

import (
	"encoding/json"
	"reflect"
)

//
// JsonDecode
//  @Description: 将`json`数据转化为结构体
//  @param content
//  @param data
//  @return interface{}
//  @return error
//
func JsonDecode(content []byte, data interface{}) (interface{}, error) {
	var err error
	dataType := reflect.TypeOf(data)
	if dataType.Kind() == reflect.Ptr {
		dataType = dataType.Elem()
	}
	instance := reflect.New(dataType)
	ptr := instance.Interface()
	if e := json.Unmarshal(content, &ptr); e != nil {
		return nil, err
	}
	return ptr, nil
}
