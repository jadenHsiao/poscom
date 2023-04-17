package template

import (
	"github.com/jadenHsiao/poscom/kernels"
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

//
//  Template
//  @Description: 模板相关结构体
//
type Template struct {
	listTemplate *ListTemplate
	kernels.Base
}

//
// NewTemplate
//  @Description: 实例化模板
//  @param config
//  @return *Template
//
func NewTemplate(config *kernels.Config) *Template {
	template := new(Template)
	template.Base = *kernels.NewBase(config)
	return template
}

//
// ListTemplate
//  @Description: 模板列表
//  @receiver template
//  @return *ListTemplate
//  @return error
//
func (template *Template) ListTemplate() (*ListTemplate, error) {
	securityCodeParams := utils.ToStrArr(
		template.Config.MemberCode,
		template.Timestamp,
		template.Config.ApiKey,
	)
	securityCode := utils.SecurityCode(securityCodeParams...)
	params := map[string]string{
		"reqTime":      template.Timestamp,
		"memberCode":   template.Base.Config.MemberCode,
		"securityCode": securityCode,
	}
	result, err := NewListTemplate().Exec(
		utils.GenerateParam(utils.ParseParam(params), nil),
	)
	template.listTemplate = result
	return template.listTemplate, err
}
