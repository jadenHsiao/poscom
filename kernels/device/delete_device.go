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

const DeleteDeviceUri = "http://api.poscom.cn:80/apisc/deldev"

//
//  DeleteDevice
//  @Description: 删除打印机数据返回结构体
//
type DeleteDevice struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

//
// NewDeleteDevice
//  @Description: 实例化删除打印机
//  @return *DeleteDevice
//
func NewDeleteDevice() *DeleteDevice {
	deleteDevice := new(DeleteDevice)
	return deleteDevice
}

//
// Exec
//  @Description: 发送请求
//  @receiver deleteDevice
//  @param params
//  @return *DeleteDevice
//  @return error
//
func (deleteDevice *DeleteDevice) Exec(params string) (*DeleteDevice, error) {
	ctx, err := utils.Send(DeleteDeviceUri, params, "POST")
	if err != nil {
		return nil, err
	}
	result, err := utils.JsonDecode(ctx, new(DeleteDevice))
	return result.(*DeleteDevice), err
}
