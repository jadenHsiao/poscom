package devicesetting

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

const SetVolumeUri = "http://api.poscom.cn:80/apisc/sendVolume"

//
//  SendVolume
//  @Description:打印机音量设置数据返回结构体
//
type SendVolume struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

//
// NewSendVolume
//  @Description: 实例化打印机音量设置
//  @return *SendVolume
//
func NewSendVolume() *SendVolume {
	sendVolume := new(SendVolume)
	return sendVolume
}

//
// Exec
//  @Description: 发起请求
//  @receiver sendVolume
//  @param params
//  @return *SendVolume
//  @return error
//
func (sendVolume *SendVolume) Exec(params string) (*SendVolume, error) {
	ctx, err := utils.Send(SetVolumeUri, params, "POST")
	if err != nil {
		return nil, err
	}
	result, err := utils.JsonDecode(ctx, new(SendVolume))
	return result.(*SendVolume), err
}
