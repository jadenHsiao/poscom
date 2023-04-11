package setting

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
import (
	"github.com/jadenHsiao/poscom/kernels"
	"github.com/jadenHsiao/poscom/utils"
)

//
//  Setting
//  @Description: 设置相关结构体
//
type Setting struct {
	kernels.Base
}

//
// NewSetting
//  @Description:
//  @param config
//  @return *Setting
//
func NewSetting(config *kernels.Config) *Setting {
	setting := new(Setting)
	setting.Base = *kernels.NewBase(config)
	return setting
}

//
// SetPushUrl
//  @Description:设置接受服务器信息推送地址
//  @receiver setting
//  @param pushUrl
//  @return *SetPushUrl
//  @return error
//
func (setting *Setting) SetPushUrl(pushUrl string) (*SetPushUrl, error) {
	securityCodeParams := utils.ToStrArr(
		setting.Config.MemberCode,
		setting.Timestamp,
		setting.Config.ApiKey,
	)
	securityCode := utils.SecurityCode(securityCodeParams...)
	params := map[string]string{
		"reqTime":      setting.Timestamp,
		"memberCode":   setting.Base.Config.MemberCode,
		"securityCode": securityCode,
		"pushUrl":      pushUrl,
	}
	result, err := NewSetPushUrl().Exec(
		utils.GenerateParam(utils.ParseParam(params), nil),
	)
	return result, err
}
