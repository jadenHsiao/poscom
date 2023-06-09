package device

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
	"net/url"
)

//
//  Device
//  @Description: 打印机设备结构体
//
type Device struct {
	list   *List
	detail *DeviceDetail
	delete *DeleteDevice
	add    *AddDevice
	edit   *EditDevice
	kernels.Base
}

//
// NewDevice
//  @Description: 实例化打印机设备相关结构体
//  @param config
//  @return *Device
//
func NewDevice(config *kernels.Config) *Device {
	device := new(Device)
	device.Base = *kernels.NewBase(config)
	return device
}

//
// DeviceList
//  @Description:查询打印机列表
//  @receiver device
//  @return *List
//  @return error
//
func (device *Device) DeviceList() (*List, error) {
	securityCodeParams := utils.ToStrArr(
		device.Config.MemberCode,
		device.Timestamp,
		device.Config.ApiKey,
	)
	securityCode := utils.SecurityCode(securityCodeParams...)
	params := map[string]string{
		"reqTime":      device.Timestamp,
		"memberCode":   device.Base.Config.MemberCode,
		"securityCode": securityCode,
	}
	result, err := NewList().Exec(
		utils.GenerateParam(utils.ParseParam(params), nil),
	)
	device.list = result
	return device.list, err
}

//
// DeviceDetail
//  @Description: 查询打印机信息
//  @receiver device
//  @param deviceID
//  @return *DeviceDetail
//  @return error
//
func (device *Device) DeviceDetail(deviceID string) (*DeviceDetail, error) {
	securityCodeParams := utils.ToStrArr(
		device.Config.MemberCode,
		device.Timestamp,
		device.Config.ApiKey,
		deviceID,
	)
	securityCode := utils.SecurityCode(securityCodeParams...)
	params := map[string]string{
		"reqTime":      device.Timestamp,
		"memberCode":   device.Base.Config.MemberCode,
		"securityCode": securityCode,
		"deviceID":     deviceID,
	}
	result, err := NewDeviceDetail().Exec(
		utils.GenerateParam(utils.ParseParam(params), nil),
	)
	device.detail = result
	return device.detail, err
}

//
// DeleteDevice
//  @Description:删除打印机
//  @receiver device
//  @param deviceID
//  @return *DeleteDevice
//  @return error
//
func (device *Device) DeleteDevice(deviceID string) (*DeleteDevice, error) {
	securityCodeParams := utils.ToStrArr(
		device.Config.MemberCode,
		device.Timestamp,
		device.Config.ApiKey,
		deviceID,
	)
	securityCode := utils.SecurityCode(securityCodeParams...)
	params := map[string]string{
		"reqTime":      device.Timestamp,
		"memberCode":   device.Base.Config.MemberCode,
		"securityCode": securityCode,
		"deviceID":     deviceID,
	}
	result, err := NewDeleteDevice().Exec(
		utils.GenerateParam(utils.ParseParam(params), nil),
	)
	device.delete = result
	return device.delete, err
}

//
// AddDevice
//  @Description: 添加打印机设备
//  @receiver device
//  @param deviceID
//  @param devName
//  @param optionalParams
//  @return *AddDevice
//  @return error
//
func (device *Device) AddDevice(deviceID string, devName string, optionalParams map[string]string) (*AddDevice, error) {
	securityCodeParams := utils.ToStrArr(
		device.Config.MemberCode,
		device.Timestamp,
		device.Config.ApiKey,
		deviceID,
	)
	securityCode := utils.SecurityCode(securityCodeParams...)
	params := map[string]string{
		"reqTime":      device.Timestamp,
		"memberCode":   device.Base.Config.MemberCode,
		"securityCode": securityCode,
		"deviceID":     deviceID,
		"devName":      url.QueryEscape(devName),
	}
	if 0 != len(optionalParams) {
		for key, val := range optionalParams {
			params[key] = val
		}
	}
	result, err := NewAddDevice().Exec(
		utils.GenerateParam(utils.ParseParam(params), nil),
	)
	device.add = result
	return device.add, err
}

//
// EditDevice
//  @Description: 编辑打印机设备
//  @receiver device
//  @param deviceID
//  @param optionalParams
//  @return *EditDevice
//  @return error
//
func (device *Device) EditDevice(deviceID string, optionalParams map[string]string) (*EditDevice, error) {
	securityCodeParams := utils.ToStrArr(
		device.Config.MemberCode,
		device.Timestamp,
		device.Config.ApiKey,
		deviceID,
	)
	securityCode := utils.SecurityCode(securityCodeParams...)
	params := map[string]string{
		"reqTime":      device.Timestamp,
		"memberCode":   device.Base.Config.MemberCode,
		"securityCode": securityCode,
		"deviceID":     deviceID,
	}
	if 0 != len(optionalParams) {
		for key, val := range optionalParams {
			if "devName" == key {
				params[key] = url.QueryEscape(val)
			} else {
				params[key] = val
			}
		}
	}
	result, err := NewEditDevice().Exec(
		utils.GenerateParam(utils.ParseParam(params), nil),
	)
	device.edit = result
	return device.edit, err
}
