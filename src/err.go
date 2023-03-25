// Package src
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
package src

//
//  PoscomError
//  @Description: 错误处理结构体
//
type PoscomError struct {
	Code      string
	Message   string
	ErrorList map[string]string
}

//
// initialize
//  @Description: 初始化
//  @receiver poscomError
//
func (poscomError *PoscomError) initialize() {
	var errorList map[string]string
	errorList = map[string]string{
		"101001": "网络请求出错",
		"101002": "数据内容有误",
	}
	poscomError.ErrorList = errorList
}

//
// NewError
//  @Description: 外部访问接口
//  @param code
//  @param message
//  @return *PoscomError
//
func NewError(code, message string) *PoscomError {
	poscomError := &PoscomError{}
	poscomError.initialize()
	poscomError.Code = code
	if 0 == len(message) {
		poscomError.Message = poscomError.ErrorList[code]
	} else {
		poscomError.Message = message
	}
	return poscomError
}
