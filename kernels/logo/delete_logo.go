package logo

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

const DeleteLogoUri = "http://api.poscom.cn:80/apisc/deleteLogo"

//
//  SetLogo
//  @Description:删除打印机`NVLogo`数据返回结构体
//
type DeleteLogo struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

//
// NewDeleteLogo
//  @Description: 实例化删除打印机`NVLogo`
//  @return *DeleteLogo
//
func NewDeleteLogo() *DeleteLogo {
	deleteLogo := new(DeleteLogo)
	return deleteLogo
}

//
// Exec
//  @Description: 发起请求
//  @receiver deleteLogo
//  @param params
//  @return *DeleteLogo
//  @return error
//
func (deleteLogo *DeleteLogo) Exec(params string) (*DeleteLogo, error) {
	ctx, err := utils.Submit(DeleteLogoUri, params, "POST")
	if err != nil {
		return nil, err
	}
	result, err := utils.JsonDecode(ctx, new(DeleteLogo))
	return result.(*DeleteLogo), err
}
