package poscom

import (
	"github.com/jadenHsiao/poscom/kernels"
	"github.com/jadenHsiao/poscom/kernels/device"
	"github.com/jadenHsiao/poscom/kernels/devicesetting"
	"github.com/jadenHsiao/poscom/kernels/group"
	"github.com/jadenHsiao/poscom/kernels/setting"
)

//  佳博打印结构体
type Gainscha struct {
	config        *kernels.Config
	group         *group.Group
	setting       *setting.Setting
	devicesetting *devicesetting.DeviceSetting
	device        *device.Device
}

//
// NewGainscha
//  @Description: 实例化佳博打印结构体函数
//  @return *Gainscha
//
func NewGainscha(cfg *kernels.Config) *Gainscha {
	return &Gainscha{
		config: cfg,
	}
}

func (gainscha *Gainscha) Group() *group.Group {
	gainscha.group = group.NewGroup(gainscha.config)
	return gainscha.group
}

func (gainscha *Gainscha) Setting() *setting.Setting {
	gainscha.setting = setting.NewSetting(gainscha.config)
	return gainscha.setting
}

func (gainscha *Gainscha) DeviceSetting() *devicesetting.DeviceSetting {
	gainscha.devicesetting = devicesetting.NewDeviceSetting(gainscha.config)
	return gainscha.devicesetting
}

func (gainscha *Gainscha) Device() *device.Device {
	gainscha.device = device.NewDevice(gainscha.config)
	return gainscha.device
}
