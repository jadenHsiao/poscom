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

const CancelPrintUri = "http://api.poscom.cn/apisc/cancelPrint"

//
//  CancelPrint
//  @Description:删除账号下指定模板数据返回结构体
//
type CancelPrint struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

//
// NewCancelPrint
//  @Description:实例化删除账号下指定模板
//  @return *CancelPrint
//
func NewCancelPrint() *CancelPrint {
	cancelPrint := new(CancelPrint)
	return cancelPrint
}

//
// Exec
//  @Description:发起请求
//  @receiver cancelPrint
//  @param params
//  @return *CancelPrint
//  @return error
//
func (cancelPrint *CancelPrint) Exec(params string) (*CancelPrint, error) {
	ctx, err := utils.Send(CancelPrintUri, params, "POST")
	if err != nil {
		return nil, err
	}
	result, err := utils.JsonDecode(ctx, new(CancelPrint))
	return result.(*CancelPrint), err
}
