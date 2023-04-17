package template

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
const ListTemplateUri = "http://api.poscom.cn:80/apisc/listTemplate"

//
//  ListTemplate
//  @Description: 模板列表返回数据结构
//
type ListTemplate struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data []Data `json:"data"`
}

//
//  Data
//  @Description: 模板列表结构体
//
type Data struct {
	Code   string              `json:"code"`
	Type   string              `json:"type"`
	Title  string              `json:"title"`
	Thumb  string              `json:"thumb"`
	Params []map[string]string `json:"params"`
}

//
// NewListTemplate
//  @Description:实例化模板列表
//  @return *ListTemplate
//
func NewListTemplate() *ListTemplate {
	listTemplate := new(ListTemplate)
	return listTemplate
}

//
// Exec
//  @Description: 发起请求
//  @receiver listTemplate
//  @param params
//  @return *ListTemplate
//  @return error
//
func (listTemplate *ListTemplate) Exec(params string) (*ListTemplate, error) {
	ctx, err := utils.Send(ListTemplateUri, params, "POST")
	if err != nil {
		return nil, err
	}
	result, err := utils.JsonDecode(ctx, new(ListTemplate))
	return result.(*ListTemplate), err
}
