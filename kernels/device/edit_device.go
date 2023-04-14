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

const EditDeviceUri = "https://api.poscom.cn/apisc/editdev"

//
//  EditDevice
//  @Description: 修改设备信息
//
type EditDevice struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

//
// NewEditDevice
//  @Description: 实例化修改设备信息
//  @return *EditDevice
//
func NewEditDevice() *EditDevice {
	editDevice := new(EditDevice)
	return editDevice
}

//
// Exec
//  @Description: 发起请求
//  @receiver editDevice
//  @param params
//  @return *EditDevice
//  @return error
//
func (editDevice *EditDevice) Exec(params string) (*EditDevice, error) {
	ctx, err := utils.Send(EditDeviceUri, params, "POST")
	if err != nil {
		return nil, err
	}
	result, err := utils.JsonDecode(ctx, new(EditDevice))
	return result.(*EditDevice), err
}
