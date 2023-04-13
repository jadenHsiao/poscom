package printer

import (
	"github.com/jadenHsiao/poscom/utils"
)

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

const GetStatusUri = "http://api.poscom.cn/apisc/getStatus"

//
//  GetStatus
//  @Description: 获取打印机状态数据返回结构体
//
type GetStatus struct {
	Code       int          `json:"code"`
	Msg        string       `json:"msg"`
	StatusList []StatusList `json:"statusList"`
}

//
//  StatusList
//  @Description: 打印机状态结构体
//
type StatusList struct {
	DeviceId string `json:"deviceId"`
	Online   int    `json:"online"`
	Status   int    `json:"status"`
	Outtime  string `json:"outtime"`
	Printnum int    `json:"printnum"`
}

//
// NewGetStatus
//  @Description: 实例化获取打印机状态
//  @return *GetStatus
//
func NewGetStatus() *GetStatus {
	getStatus := new(GetStatus)
	return getStatus
}

//
// Exec
//  @Description: 发起请求
//  @receiver getStatus
//  @param params
//  @return *GetStatus
//  @return error
//
func (getStatus *GetStatus) Exec(params string) (*GetStatus, error) {
	ctx, err := utils.Send(GetStatusUri, params, "GET")
	if err != nil {
		return nil, err
	}
	result, err := utils.JsonDecode(ctx, new(GetStatus))
	return result.(*GetStatus), err
}
