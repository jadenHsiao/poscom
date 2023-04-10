package devicesetting

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
	"strconv"
)

//
//  DeviceSetting
//  @Description:设置结构体
//
type DeviceSetting struct {
	setNet        *SetNet
	sendVolume    *SendVolume
	sendVoiceType *SendVoiceType
	kernels.Base
}

//
// NewDeviceSetting
//  @Description: 实例化设置相关结构体
//  @param config
//  @return *Setting
//
func NewDeviceSetting(config *kernels.Config) *DeviceSetting {
	deviceSetting := new(DeviceSetting)
	deviceSetting.Base = *kernels.NewBase(config)
	return deviceSetting
}

//
// SetNet
//  @Description: 设置网络链接方式
//  @receiver deviceSetting
//  @param deviceID
//  @param netType
//  @return *SetNet
//  @return error
//
func (deviceSetting *DeviceSetting) SetNet(deviceID string, netType string) (*SetNet, error) {
	securityCodeParams := utils.ToStrArr(
		deviceSetting.Config.MemberCode,
		deviceSetting.Timestamp,
		deviceSetting.Config.ApiKey,
		deviceID,
	)
	securityCode := utils.SecurityCode(securityCodeParams...)
	params := map[string]string{
		"reqTime":      deviceSetting.Timestamp,
		"memberCode":   deviceSetting.Base.Config.MemberCode,
		"securityCode": securityCode,
		"deviceID":     deviceID,
		"netType":      netType,
	}
	result, err := NewSetNet().Exec(
		utils.GenerateParam(utils.ParseParam(params), nil),
	)
	return result, err
}

//
// SendVolume
//  @Description:打印机音量设置
//  @receiver deviceSetting
//  @param deviceID
//  @param volume
//  @return *SendVolume
//  @return error
//
func (deviceSetting *DeviceSetting) SendVolume(deviceID string, volume int) (*SendVolume, error) {
	securityCodeParams := utils.ToStrArr(
		deviceSetting.Config.MemberCode,
		deviceSetting.Timestamp,
		deviceSetting.Config.ApiKey,
		deviceID,
	)
	securityCode := utils.SecurityCode(securityCodeParams...)
	params := map[string]string{
		"reqTime":      deviceSetting.Timestamp,
		"memberCode":   deviceSetting.Base.Config.MemberCode,
		"securityCode": securityCode,
		"deviceID":     deviceID,
		"volume":       strconv.Itoa(volume),
	}
	result, err := NewSendVolume().Exec(
		utils.GenerateParam(utils.ParseParam(params), nil),
	)
	return result, err
}

//
// SendVoiceType
//  @Description:打印机切换播报类型
//  @receiver deviceSetting
//  @param deviceID
//  @param voiceType
//  @return *SendVoiceType
//  @return error
//
func (deviceSetting *DeviceSetting) SendVoiceType(deviceID string, voiceType int) (*SendVoiceType, error) {
	securityCodeParams := utils.ToStrArr(
		deviceSetting.Config.MemberCode,
		deviceSetting.Timestamp,
		deviceSetting.Config.ApiKey,
		deviceID,
	)
	securityCode := utils.SecurityCode(securityCodeParams...)
	params := map[string]string{
		"reqTime":      deviceSetting.Timestamp,
		"memberCode":   deviceSetting.Base.Config.MemberCode,
		"securityCode": securityCode,
		"deviceID":     deviceID,
		"voiceType":    strconv.Itoa(voiceType),
	}
	result, err := NewSendVoiceType().Exec(
		utils.GenerateParam(utils.ParseParam(params), nil),
	)
	return result, err
}
