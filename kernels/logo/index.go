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

import (
	"github.com/jadenHsiao/poscom/kernels"
	"github.com/jadenHsiao/poscom/utils"
)

//
//  Logo
//  @Description: `NVLogo`设置结构体
//
type Logo struct {
	setLogo    *SetLogo
	deleteLogo *DeleteLogo
	kernels.Base
}

//
// NewLogo
//  @Description: 实例化`logo`相关结构体
//  @param config
//  @return *Logo
//
func NewLogo(config *kernels.Config) *Logo {
	logo := new(Logo)
	logo.Base = *kernels.NewBase(config)
	return logo
}

//
// SetLogo
//  @Description:设置`logo`
//  @receiver logo
//  @param deviceID
//  @param imgUrl
//  @return *SetLogo
//  @return error
//
func (logo *Logo) SetLogo(deviceID string, imgUrl string) (*SetLogo, error) {
	securityCodeParams := utils.ToStrArr(
		logo.Config.MemberCode,
		logo.Timestamp,
		logo.Config.ApiKey,
		deviceID,
	)
	securityCode := utils.SecurityCode(securityCodeParams...)
	params := map[string]string{
		"reqTime":      logo.Timestamp,
		"memberCode":   logo.Base.Config.MemberCode,
		"securityCode": securityCode,
		"deviceID":     deviceID,
		"imgUrl":       imgUrl,
	}
	result, err := NewSetLogo().Exec(
		utils.GenerateParam(utils.ParseParam(params), nil),
	)
	return result, err
}

//
// DeleteLogo
//  @Description: 删除打印机`NVLogo`
//  @receiver logo
//  @param deviceID
//  @return *DeleteLogo
//  @return error
//
func (logo *Logo) DeleteLogo(deviceID string) (*DeleteLogo, error) {
	securityCodeParams := utils.ToStrArr(
		logo.Config.MemberCode,
		logo.Timestamp,
		logo.Config.ApiKey,
		deviceID,
	)
	securityCode := utils.SecurityCode(securityCodeParams...)
	params := map[string]string{
		"reqTime":      logo.Timestamp,
		"memberCode":   logo.Base.Config.MemberCode,
		"securityCode": securityCode,
		"deviceID":     deviceID,
	}
	result, err := NewDeleteLogo().Exec(
		utils.GenerateParam(utils.ParseParam(params), nil),
	)
	return result, err
}
