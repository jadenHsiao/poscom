package printer

import "github.com/jadenHsiao/poscom/utils"

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

const ListExceptionUri = "http://api.poscom.cn/apisc/listException"

//
//  ListException
//  @Description:查询打印异常信息
//
type ListException struct {
	Code   int      `json:"code"`
	Msg    string   `json:"msg"`
	Begin  string   `json:"begin"`
	End    string   `json:"end"`
	ExList []ExList `json:"exList"`
}

//
//  ExList
//  @Description: 异常列表
//
type ExList struct {
	MsgNo   string `json:"msgNo"`
	MsgTime string `json:"msgTime"`
	Detail  string `json:"detail"`
}

//
// NewListException
//  @Description: 实例化查询打印异常信息
//  @return *ListException
//
func NewListException() *ListException {
	listException := new(ListException)
	return listException
}

//
// Exec
//  @Description: 发起请求
//  @receiver listException
//  @param params
//  @return *ListException
//  @return error
//
func (listException *ListException) Exec(params string) (*ListException, error) {
	ctx, err := utils.Send(ListExceptionUri, params, "GET")
	if err != nil {
		return nil, err
	}
	result, err := utils.JsonDecode(ctx, new(ListException))
	return result.(*ListException), err
}
