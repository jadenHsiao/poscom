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

const SetNetUri = "http://api.poscom.cn:80/apisc/setNet"

//
//  SetNet
//  @Description:网络链接方式设置数据返回结构体
//
type SetNet struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

//
// NewSetNet
//  @Description: 实例化网络链接设置方式
//  @return *SetNet
//
func NewSetNet() *SetNet {
	setNet := new(SetNet)
	return setNet
}

//
// Exec
//  @Description: 发起请求
//  @receiver setNet
//  @param params
//  @return *SetNet
//  @return error
//
func (setNet *SetNet) Exec(params string) (*SetNet, error) {
	ctx, err := utils.Submit(SetNetUri, params, "POST")
	if err != nil {
		return nil, err
	}
	result, err := utils.JsonDecode(ctx, new(SetNet))
	return result.(*SetNet), err
}
