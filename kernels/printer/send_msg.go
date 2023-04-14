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
const SendMsgUri = "http://api.poscom.cn/apisc/sendMsg"

//
//  SendMsg
//  @Description: 发送数据到打印机
//
type SendMsg struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

//
// NewSendMsg
//  @Description: 实例化发送数据到打印机
//  @return *SendMsg
//
func NewSendMsg() *SendMsg {
	sendMsg := new(SendMsg)
	return sendMsg
}

//
// Exec
//  @Description: 发起请求
//  @receiver sendMsg
//  @param params
//  @return *SendMsg
//  @return error
//
func (sendMsg *SendMsg) Exec(params string) (*SendMsg, error) {
	ctx, err := utils.Send(SendMsgUri, params, "POST")
	if err != nil {
		return nil, err
	}
	result, err := utils.JsonDecode(ctx, new(SendMsg))
	return result.(*SendMsg), err
}
