// Package utils
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
package utils

import (
	"reflect"
	"strconv"
)

//
// ToStrArr
//  @Description: 参数数据类型转化为字符串
//  @param args
//  @return result
//
func ToStrArr(args ...interface{}) (result []string) {
	result = make([]string, len(args))
	for key, val := range args {
		dataType := reflect.TypeOf(val).Kind().String()
		switch dataType {
		case "string":
		case "int":
			val = strconv.Itoa(val.(int))
		case "int64":
			val = strconv.FormatInt(val.(int64), 10)
		case "float64":
			val = strconv.FormatFloat(val.(float64), 'E', -1, 64)
		}
		result[key] = val.(string)
	}
	return result
}
