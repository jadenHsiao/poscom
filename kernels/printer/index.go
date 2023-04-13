package printer

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
//  Printer
//  @Description:打印相关结构体
//
type Printer struct {
	queryState    *QueryState
	listException *ListException
	getStatus     *GetStatus
	cancelPrint   *CancelPrint
	kernels.Base
}

//
// NewPrinter
//  @Description: 实例化打印相关结构体
//  @param config
//  @return *Printer
//
func NewPrinter(config *kernels.Config) *Printer {
	printer := new(Printer)
	printer.Base = *kernels.NewBase(config)
	return printer
}

//
// QueryState
//  @Description:查询打印任务状态
//  @receiver printer
//  @param msgNo
//  @return *QueryState
//  @return error
//
func (printer *Printer) QueryState(msgNo string) (*QueryState, error) {
	securityCodeParams := utils.ToStrArr(
		printer.Config.MemberCode,
		printer.Timestamp,
		printer.Config.ApiKey,
		msgNo,
	)
	securityCode := utils.SecurityCode(securityCodeParams...)
	params := map[string]string{
		"reqTime":      printer.Timestamp,
		"memberCode":   printer.Base.Config.MemberCode,
		"securityCode": securityCode,
		"msgNo":        msgNo,
	}
	result, err := NewQueryState().Exec(
		utils.GenerateParam(utils.ParseParam(params), nil),
	)
	printer.queryState = result
	return printer.queryState, err
}

//
// ListException
//  @Description: 查询打印异常信息
//  @receiver printer
//  @param deviceID
//  @param start
//  @param end
//  @return *ListException
//  @return error
//
func (printer *Printer) ListException(deviceID string, start string, end string) (*ListException, error) {
	securityCodeParams := utils.ToStrArr(
		printer.Config.MemberCode,
		printer.Timestamp,
		printer.Config.ApiKey,
	)
	securityCode := utils.SecurityCode(securityCodeParams...)
	params := map[string]string{
		"reqTime":      printer.Timestamp,
		"memberCode":   printer.Base.Config.MemberCode,
		"securityCode": securityCode,
		"deviceID":     deviceID,
	}
	if 0 != len(start) {
		params["start"] = start
	}
	if 0 != len(end) {
		params["end"] = end
	}
	result, err := NewListException().Exec(
		utils.GenerateParam(utils.ParseParam(params), nil),
	)
	printer.listException = result
	return printer.listException, err
}

//
// GetStatus
//  @Description: 获取打印机状态
//  @receiver printer
//  @param deviceID
//  @return *GetStatus
//  @return error
//
func (printer *Printer) GetStatus(deviceID string) (*GetStatus, error) {
	securityCodeParams := utils.ToStrArr(
		printer.Config.MemberCode,
		printer.Timestamp,
		printer.Config.ApiKey,
	)
	securityCode := utils.SecurityCode(securityCodeParams...)
	params := map[string]string{
		"reqTime":      printer.Timestamp,
		"memberCode":   printer.Base.Config.MemberCode,
		"securityCode": securityCode,
	}
	if 0 != len(deviceID) {
		params["deviceID"] = deviceID
	}
	result, err := NewGetStatus().Exec(
		utils.GenerateParam(utils.ParseParam(params), nil),
	)
	printer.getStatus = result
	return printer.getStatus, err
}

//
// CancelPrint
//  @Description:删除账号下指定模板
//  @receiver printer
//  @param deviceID
//  @param all
//  @return *CancelPrint
//  @return error
//
func (printer *Printer) CancelPrint(deviceID string, all string) (*CancelPrint, error) {
	securityCodeParams := utils.ToStrArr(
		printer.Config.MemberCode,
		printer.Timestamp,
		printer.Config.ApiKey,
		deviceID,
	)
	securityCode := utils.SecurityCode(securityCodeParams...)
	params := map[string]string{
		"reqTime":      printer.Timestamp,
		"memberCode":   printer.Base.Config.MemberCode,
		"securityCode": securityCode,
		"deviceID":     deviceID,
		"all":          all,
	}
	result, err := NewCancelPrint().Exec(
		utils.GenerateParam(utils.ParseParam(params), nil),
	)
	printer.cancelPrint = result
	return printer.cancelPrint, err
}
