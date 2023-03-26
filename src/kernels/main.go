// Package kernels
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
package kernels

import (
	"fmt"
	"github.com/jadenHsiao/poscom/src"
	"net/url"
	"strconv"
	"strings"
)

//
//  Gainscha
//  @Description: 佳博打印结构体
//
type Gainscha struct {
	Api          *src.Api
	ApiSecretKey string
	MemberCode   string
}

//
// initialize
//  @Description: 初始化佳博打印结构体
//  @receiver gainscha
//
func (gainscha *Gainscha) initialize() {
	gainscha.Api = new(src.Api)
	gainscha.Api.Initialize()
}

//
// generateParam
//  @Description: 创建请求参数
//  @receiver gainscha
//  @param reqTime
//  @param securityCode
//  @param otherParams
//  @return string
//
func (gainscha *Gainscha) generateParam(reqTime int64, securityCode string, otherParams map[string]string) string {
	reqTimeString := strconv.FormatInt(reqTime, 10)
	paramList := []string{
		"reqTime=" + reqTimeString,
		"securityCode=" + securityCode,
		"memberCode=" + gainscha.MemberCode,
	}
	for key, val := range otherParams {
		item := fmt.Sprintf("%v=%v", key, val)
		paramList = append(paramList, item)
	}
	return strings.Join(paramList, "&")
}

//
// securityCode
//  @Description: 根据官方规则创建秘钥
//  @receiver gainscha
//  @param args
//  @return string
//
func (gainscha *Gainscha) securityCode(args ...string) string {
	var beforeStr string
	for _, val := range args {
		beforeStr += val
	}
	return src.Md5(beforeStr)
}

//
// ListDevice
//  @Description: 查询打印机列表
// 	@See:https://dev.poscom.cn/openapi?listDevice
//  @receiver gainscha
//  @return result
//  @return err
//
func (gainscha *Gainscha) ListDevice() (result map[string]interface{}, err *src.PoscomError) {
	gainscha.initialize()
	reqTime := src.Time()
	securityCodeParams := src.ArgsType2String(
		gainscha.MemberCode,
		reqTime,
		gainscha.ApiSecretKey,
	)
	securityCode := gainscha.securityCode(securityCodeParams...)
	params := gainscha.generateParam(reqTime, securityCode, nil)
	request := new(Request)
	request.Method = "GET"
	target := gainscha.Api.List["ListDevice"]
	return request.Send(fmt.Sprintf("%v?%v", target, params))
}

//
// Device
//  @Description:查询打印机信息
//	@See:https://dev.poscom.cn/openapi?device
//  @receiver gainscha
//  @param deviceID
//  @return result
//  @return err
//
func (gainscha *Gainscha) Device(deviceID string) (result map[string]interface{}, err *src.PoscomError) {
	gainscha.initialize()
	reqTime := src.Time()
	securityCodeParams := src.ArgsType2String(
		gainscha.MemberCode,
		reqTime,
		gainscha.ApiSecretKey,
		deviceID,
	)
	securityCode := gainscha.securityCode(securityCodeParams...)
	otherParams := map[string]string{
		"deviceID": deviceID,
	}
	params := gainscha.generateParam(reqTime, securityCode, otherParams)
	request := new(Request)
	request.Method = "GET"
	target := gainscha.Api.List["Device"]
	return request.Send(fmt.Sprintf("%v?%v", target, params))
}

//
// GetStatus
//  @Description:获取打印机状态
//	@See:https://dev.poscom.cn/openapi?getStatus
//  @receiver gainscha
//  @param deviceID
//  @return result
//  @return err
//
func (gainscha *Gainscha) GetStatus(deviceID string) (result map[string]interface{}, err *src.PoscomError) {
	gainscha.initialize()
	reqTime := src.Time()
	securityCodeParams := src.ArgsType2String(
		gainscha.MemberCode,
		reqTime,
		gainscha.ApiSecretKey,
	)
	securityCode := gainscha.securityCode(securityCodeParams...)
	otherParams := map[string]string{
		"deviceID": deviceID,
	}
	params := gainscha.generateParam(reqTime, securityCode, otherParams)
	request := new(Request)
	request.Method = "POST"
	target := gainscha.Api.List["GetStatus"]
	return request.Send(fmt.Sprintf("%v?%v", target, params))
}

//
// DelDev
//  @Description:删除打印机
//	@See:https://dev.poscom.cn/openapi?deldev
//  @receiver gainscha
//  @param deviceID
//  @return result
//  @return err
//
func (gainscha *Gainscha) DelDev(deviceID string) (result map[string]interface{}, err *src.PoscomError) {
	gainscha.initialize()
	reqTime := src.Time()
	securityCodeParams := src.ArgsType2String(
		gainscha.MemberCode,
		reqTime,
		gainscha.ApiSecretKey,
		deviceID,
	)
	securityCode := gainscha.securityCode(securityCodeParams...)
	otherParams := map[string]string{
		"deviceID": deviceID,
	}
	params := gainscha.generateParam(reqTime, securityCode, otherParams)
	request := new(Request)
	request.Method = "POST"
	target := gainscha.Api.List["DelDev"]
	return request.Send(fmt.Sprintf("%v?%v", target, params))
}

//
// Group
//  @Description: 获取全部分组
//  @receiver gainscha
//  @return result
//  @return err
//
func (gainscha *Gainscha) Group() (result map[string]interface{}, err *src.PoscomError) {
	gainscha.initialize()
	reqTime := src.Time()
	securityCodeParams := src.ArgsType2String(
		gainscha.MemberCode,
		reqTime,
		gainscha.ApiSecretKey,
	)
	securityCode := gainscha.securityCode(securityCodeParams...)
	params := gainscha.generateParam(reqTime, securityCode, nil)
	request := new(Request)
	request.Method = "GET"
	target := gainscha.Api.List["Group"]
	return request.Send(fmt.Sprintf("%v?%v", target, params))
}

//
// ListTemplate
//  @Description:模板列表
//	@See:https://dev.poscom.cn/openapi?templetList
//  @receiver gainscha
//  @return result
//  @return err
//
func (gainscha *Gainscha) ListTemplate() (result map[string]interface{}, err *src.PoscomError) {
	gainscha.initialize()
	reqTime := src.Time()
	securityCodeParams := src.ArgsType2String(
		gainscha.MemberCode,
		reqTime,
		gainscha.ApiSecretKey,
	)
	securityCode := gainscha.securityCode(securityCodeParams...)
	params := gainscha.generateParam(reqTime, securityCode, nil)
	request := new(Request)
	request.Method = "POST"
	target := gainscha.Api.List["ListTemplate"]
	return request.Send(fmt.Sprintf("%v?%v", target, params))
}

//
// AddGroup
//  @Description:添加（打印机）分组
//  @receiver gainscha
//  @param grpName
//  @return result
//  @return err
//
func (gainscha *Gainscha) AddGroup(grpName string) (result map[string]interface{}, err *src.PoscomError) {
	gainscha.initialize()
	reqTime := src.Time()
	securityCodeParams := src.ArgsType2String(
		gainscha.MemberCode,
		reqTime,
		gainscha.ApiSecretKey,
	)
	securityCode := gainscha.securityCode(securityCodeParams...)
	otherParams := map[string]string{
		"grpName": url.QueryEscape(grpName),
	}
	params := gainscha.generateParam(reqTime, securityCode, otherParams)
	request := new(Request)
	request.Method = "POST"
	target := gainscha.Api.List["AddGroup"]
	return request.Send(fmt.Sprintf("%v?%v", target, params))
}

//
// ListException
//  @Description: 查询打印异常信息
//  @receiver gainscha
//  @param deviceNo
//  @param start
//  @param end
//  @return result
//  @return err
//
func (gainscha *Gainscha) ListException(deviceNo string, start string, end string) (result map[string]interface{}, err *src.PoscomError) {
	gainscha.initialize()
	reqTime := src.Time()
	securityCodeParams := src.ArgsType2String(
		gainscha.MemberCode,
		reqTime,
		gainscha.ApiSecretKey,
	)
	securityCode := gainscha.securityCode(securityCodeParams...)
	otherParams := map[string]string{
		"imsi":  deviceNo,
		"start": start,
		"end":   end,
	}
	params := gainscha.generateParam(reqTime, securityCode, otherParams)
	request := new(Request)
	request.Method = "GET"
	target := gainscha.Api.List["ListException"]
	return request.Send(fmt.Sprintf("%v?%v", target, params))
}

//
// EditGroup
//  @Description:修改（打印机）分组名称
//  @receiver gainscha
//  @param grpId
//  @param grpName
//  @return result
//  @return err
//
func (gainscha *Gainscha) EditGroup(grpId string, grpName string) (result map[string]interface{}, err *src.PoscomError) {
	gainscha.initialize()
	reqTime := src.Time()
	securityCodeParams := src.ArgsType2String(
		gainscha.MemberCode,
		reqTime,
		gainscha.ApiSecretKey,
		grpId,
	)
	securityCode := gainscha.securityCode(securityCodeParams...)
	otherParams := map[string]string{
		"grpID":   grpId,
		"grpName": url.QueryEscape(grpName),
	}
	params := gainscha.generateParam(reqTime, securityCode, otherParams)
	request := new(Request)
	request.Method = "POST"
	target := gainscha.Api.List["EditGroup"]
	return request.Send(fmt.Sprintf("%v?%v", target, params))
}

//
// DelGroup
//  @Description:删除（打印机）分组名称
//  @receiver gainscha
//  @param grpId
//  @return result
//  @return err
//
func (gainscha *Gainscha) DelGroup(grpId string) (result map[string]interface{}, err *src.PoscomError) {
	gainscha.initialize()
	reqTime := src.Time()
	securityCodeParams := src.ArgsType2String(
		gainscha.MemberCode,
		reqTime,
		gainscha.ApiSecretKey,
		grpId,
	)
	securityCode := gainscha.securityCode(securityCodeParams...)
	otherParams := map[string]string{
		"grpID": grpId,
	}
	params := gainscha.generateParam(reqTime, securityCode, otherParams)
	request := new(Request)
	request.Method = "POST"
	target := gainscha.Api.List["DelGroup"]
	return request.Send(fmt.Sprintf("%v?%v", target, params))
}

//
// QueryState
//  @Description:查询打印任务状态
//	@See:https://dev.poscom.cn/openapi?queryState
//  @receiver gainscha
//  @param msgNo
//  @return result
//  @return err
//
func (gainscha *Gainscha) QueryState(msgNo string) (result map[string]interface{}, err *src.PoscomError) {
	gainscha.initialize()
	reqTime := src.Time()
	securityCodeParams := src.ArgsType2String(
		gainscha.MemberCode,
		reqTime,
		gainscha.ApiSecretKey,
		msgNo,
	)
	securityCode := gainscha.securityCode(securityCodeParams...)
	otherParams := map[string]string{
		"msgNo": msgNo,
	}
	params := gainscha.generateParam(reqTime, securityCode, otherParams)
	request := new(Request)
	request.Method = "GET"
	target := gainscha.Api.List["QueryState"]
	return request.Send(fmt.Sprintf("%v?%v", target, params))
}

//
// DeleteLogo
//  @Description:删除NVLogo（票据机）
// 	@See:https://dev.poscom.cn/openapi?deleteLogo
//  @receiver gainscha
//  @param deviceID
//  @return result
//  @return err
//
func (gainscha *Gainscha) DeleteLogo(deviceID string) (result map[string]interface{}, err *src.PoscomError) {
	gainscha.initialize()
	reqTime := src.Time()
	securityCodeParams := src.ArgsType2String(
		gainscha.MemberCode,
		reqTime,
		gainscha.ApiSecretKey,
		deviceID,
	)
	securityCode := gainscha.securityCode(securityCodeParams...)
	otherParams := map[string]string{
		"deviceID": deviceID,
	}
	params := gainscha.generateParam(reqTime, securityCode, otherParams)
	request := new(Request)
	request.Method = "POST"
	target := gainscha.Api.List["DeleteLogo"]
	return request.Send(fmt.Sprintf("%v?%v", target, params))
}
