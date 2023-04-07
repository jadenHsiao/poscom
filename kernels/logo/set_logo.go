package logo

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

import "github.com/jadenHsiao/poscom/utils"

const SetLogoUri = "http://api.poscom.cn:80/apisc/setLogo"

//
//  SetLogo
//  @Description:设置打印机`NVLogo`数据返回结构体
//
type SetLogo struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

//
// NewSetLogo
//  @Description: 实例化设置`logo`
//  @return *SetLogo
//
func NewSetLogo() *SetLogo {
	setLogo := new(SetLogo)
	return setLogo
}

//
// Exec
//  @Description: 发起请求
//  @receiver setLogo
//  @param params
//  @return *SetLogo
//  @return error
//
func (setLogo *SetLogo) Exec(params string) (*SetLogo, error) {
	ctx, err := utils.Submit(SetLogoUri, params, "POST")
	if err != nil {
		return nil, err
	}
	result, err := utils.JsonDecode(ctx, new(SetLogo))
	return result.(*SetLogo), err
}
