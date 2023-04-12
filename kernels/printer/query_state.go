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

const QueryStateUri = "http://api.poscom.cn/apisc/queryState"

//
//  QueryState
//  @Description:查询打印任务状态数据返回结构体
//
type QueryState struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

//
// NewQueryState
//  @Description: 实例化查询打印任务状态
//  @return *QueryState
//
func NewQueryState() *QueryState {
	queryState := new(QueryState)
	return queryState
}

//
// Exec
//  @Description: 发起请求
//  @receiver queryState
//  @param params
//  @return *QueryState
//  @return error
//
func (queryState *QueryState) Exec(params string) (*QueryState, error) {
	ctx, err := utils.Send(QueryStateUri, params, "POST")
	if err != nil {
		return nil, err
	}
	result, err := utils.JsonDecode(ctx, new(QueryState))
	return result.(*QueryState), err
}
