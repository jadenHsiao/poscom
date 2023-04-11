package device

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

const DeviceUri = "http://api.poscom.cn:80/apisc/device"

//
//  DeviceDetail
//  @Description:查询打印机信息数据返回结构体
//
type DeviceDetail struct {
	Code    int     `json:"code"`
	Msg     string  `json:"msg"`
	DevInfo DevInfo `json:"devInfo"`
}

//
//  DevInfo
//  @Description: 设备信息结构体
//
type DevInfo struct {
	DeviceID string `json:"deviceID"`
	DevName  string `json:"devName"`
	Mphone   string `json:"mPhone"`
	Nphone   string `json:"nPhone"`
	Cutting  string `json:"cutting"`
}

//
// NewDeviceDetail
//  @Description: 实例化查询打印机信息
//  @return *DeviceDetail
//
func NewDeviceDetail() *DeviceDetail {
	deviceDetail := new(DeviceDetail)
	return deviceDetail
}

//
// Exec
//  @Description: 发起请求
//  @receiver deviceDetail
//  @param params
//  @return *DeviceDetail
//  @return error
//
func (deviceDetail *DeviceDetail) Exec(params string) (*DeviceDetail, error) {
	ctx, err := utils.Send(DeviceUri, params, "POST")
	if err != nil {
		return nil, err
	}
	result, err := utils.JsonDecode(ctx, new(DeviceDetail))
	return result.(*DeviceDetail), err
}
