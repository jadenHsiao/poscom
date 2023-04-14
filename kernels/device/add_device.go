package device

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

const AddDeviceUri = "http://api.poscom.cn:80/apisc/adddev"

//
//  AddDevice
//  @Description: 添加打印机设备数据返回结构体
//
type AddDevice struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

//
// NewAddDevice
//  @Description: 实例化添加打印机
//  @return *AddDevice
//
func NewAddDevice() *AddDevice {
	addDevice := new(AddDevice)
	return addDevice
}

//
// Exec
//  @Description: 发起请求
//  @receiver addDevice
//  @param params
//  @return *AddDevice
//  @return error
//
func (addDevice *AddDevice) Exec(params string) (*AddDevice, error) {
	ctx, err := utils.Send(AddDeviceUri, params, "POST")
	if err != nil {
		return nil, err
	}
	result, err := utils.JsonDecode(ctx, new(AddDevice))
	return result.(*AddDevice), err
}
