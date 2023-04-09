package devicesetting

import (
	"github.com/jadenHsiao/poscom/kernels"
	"github.com/jadenHsiao/poscom/utils"
)

//
//  DeviceSetting
//  @Description:设置结构体
//
type DeviceSetting struct {
	setNet *SetNet
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
