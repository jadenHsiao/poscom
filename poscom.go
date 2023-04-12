package poscom

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
func NewGainscha(apiKey string, memberCode string) *Gainscha {
	cfg := kernels.Config{ApiKey: apiKey, MemberCode: memberCode}
	return &Gainscha{
		config: &cfg,
	}
}

//
// Group
//  @Description: 打印机分组
//  @receiver gainscha
//  @return *group.Group
//
func (gainscha *Gainscha) Group() *group.Group {
	gainscha.group = group.NewGroup(gainscha.config)
	return gainscha.group
}

//
// Setting
//  @Description: 设置
//  @receiver gainscha
//  @return *setting.Setting
//
func (gainscha *Gainscha) Setting() *setting.Setting {
	gainscha.setting = setting.NewSetting(gainscha.config)
	return gainscha.setting
}

//
// DeviceSetting
//  @Description: 打印机设置
//  @receiver gainscha
//  @return *devicesetting.DeviceSetting
//
func (gainscha *Gainscha) DeviceSetting() *devicesetting.DeviceSetting {
	gainscha.devicesetting = devicesetting.NewDeviceSetting(gainscha.config)
	return gainscha.devicesetting
}

//
// Device
//  @Description: 打印机
//  @receiver gainscha
//  @return *device.Device
//
func (gainscha *Gainscha) Device() *device.Device {
	gainscha.device = device.NewDevice(gainscha.config)
	return gainscha.device
}
