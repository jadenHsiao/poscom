package setting

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

const SetPushUrlUri = "http://api.poscom.cn:80/apisc/setPushUrl"

//
//  SetPushUrl
//  @Description:设置接受服务器信息推送地址数据返回结构体
//
type SetPushUrl struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

//
// NewSetPushUrl
//  @Description: 实例化设置接受服务器信息推送地址
//  @return *SetPushUrl
//
func NewSetPushUrl() *SetPushUrl {
	setPushUrl := new(SetPushUrl)
	return setPushUrl
}

//
// Exec
//  @Description: 发起请求
//  @receiver setPushUrl
//  @param params
//  @return *SetPushUrl
//  @return error
//
func (setPushUrl *SetPushUrl) Exec(params string) (*SetPushUrl, error) {
	ctx, err := utils.Send(SetPushUrlUri, params, "POST")
	if err != nil {
		return nil, err
	}
	result, err := utils.JsonDecode(ctx, new(SetPushUrl))
	return result.(*SetPushUrl), err
}
