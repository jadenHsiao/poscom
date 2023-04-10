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

const SetVoiceTypeUri = "https://api.poscom.cn:80/apisc/setVoiceType"

//
//  SendVoiceType
//  @Description:打印机切换播报类型数据返回结构体
//
type SendVoiceType struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

//
// NewSendVoiceType
//  @Description: 实例化打印机切换播报类型
//  @return *SendVoiceType
//
func NewSendVoiceType() *SendVoiceType {
	sendVoiceType := new(SendVoiceType)
	return sendVoiceType
}

//
// Exec
//  @Description: 发起请求
//  @receiver sendVoiceType
//  @param params
//  @return *SendVoiceType
//  @return error
//
func (sendVoiceType *SendVoiceType) Exec(params string) (*SendVoiceType, error) {
	ctx, err := utils.Submit(SetVoiceTypeUri, params, "POST")
	if err != nil {
		return nil, err
	}
	result, err := utils.JsonDecode(ctx, new(SendVoiceType))
	return result.(*SendVoiceType), err
}
