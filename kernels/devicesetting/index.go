package devicesetting

import "github.com/jadenHsiao/poscom/kernels"

//
//  DeviceSetting
//  @Description:设置结构体
//
type DeviceSetting struct {
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
