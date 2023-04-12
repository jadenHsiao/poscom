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
	queryState *QueryState
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
