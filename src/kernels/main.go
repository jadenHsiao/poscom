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
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/jadenHsiao/poscom/src"
)

// Gainscha
// @Description: 佳博打印结构体
type Gainscha struct {
	Api          *src.Api
	ApiSecretKey string
	MemberCode   string
	request      *Request
	err          *src.PoscomError
}

// initialize
//
//	@Description: 初始化佳博打印结构体
//	@receiver gainscha
func (gainscha *Gainscha) initialize() {
	gainscha.Api = new(src.Api)
	gainscha.Api.Initialize()
	gainscha.request = new(Request)
}

// generateParam
//
//	@Description: 创建请求参数
//	@receiver gainscha
//	@param reqTime
//	@param securityCode
//	@param otherParams
//	@return string
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

// securityCode
//
//	@Description: 根据官方规则创建秘钥
//	@receiver gainscha
//	@param args
//	@return string
func (gainscha *Gainscha) securityCode(args ...string) string {
	var beforeStr string
	for _, val := range args {
		beforeStr += val
	}
	return src.Md5(beforeStr)
}

//
// ToMap
//  @Description: 返回结构体
//  @receiver gainscha
//  @return result
//  @return err
//
func (gainscha *Gainscha) ToMap() (result map[string]interface{}, err *src.PoscomError) {
	if nil != gainscha.err {
		return nil, gainscha.err
	}
	body := gainscha.request.Response
	errContent := json.Unmarshal(body, &result)
	if nil != errContent {
		gainscha.err = src.NewError("101002", errContent.Error())
		return nil, gainscha.err
	}
	return result, nil
}

//
// ToString
//  @Description: 返回字符串
//  @receiver gainscha
//  @return result
//  @return err
//
func (gainscha *Gainscha) ToString() (result string, err *src.PoscomError) {
	if nil != gainscha.err {
		return "", gainscha.err
	}
	body := gainscha.request.Response
	return string(body), nil
}

//
// ToByte
//  @Description: 返回`byte`
//  @receiver gainscha
//  @return result
//  @return err
//
func (gainscha *Gainscha) ToByte() (result []byte, err *src.PoscomError) {
	if nil != gainscha.err {
		return nil, gainscha.err
	}
	return gainscha.request.Response, nil
}

//
// ListDevice
// @Description: 查询打印机列表
// @See:https://dev.poscom.cn/openapi?listDevice
// @receiver gainscha
// @return *Gainscha
//
func (gainscha *Gainscha) ListDevice() *Gainscha {
	gainscha.initialize()
	reqTime := src.Time()
	securityCodeParams := src.ArgsType2String(
		gainscha.MemberCode,
		reqTime,
		gainscha.ApiSecretKey,
	)
	securityCode := gainscha.securityCode(securityCodeParams...)
	params := gainscha.generateParam(reqTime, securityCode, nil)
	request := gainscha.request
	request.Method = "GET"
	target := gainscha.Api.List["ListDevice"]
	resp := request.Send(fmt.Sprintf("%v?%v", target, params))
	if nil != resp {
		gainscha.err = resp
	}
	return gainscha
}

//
// Device
//	@Description:查询打印机信息
//	@See:https://dev.poscom.cn/openapi?device
//  @receiver gainscha
//  @param deviceID
//  @return *Gainscha
//
func (gainscha *Gainscha) Device(deviceID string) *Gainscha {
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
	request := gainscha.request
	request.Method = "GET"
	target := gainscha.Api.List["Device"]
	resp := request.Send(fmt.Sprintf("%v?%v", target, params))
	if nil != resp {
		gainscha.err = resp
	}
	return gainscha
}

//
// GetStatus
//	@Description:获取打印机状态
//	@See:https://dev.poscom.cn/openapi?getStatus
//  @receiver gainscha
//  @param deviceID
//  @return *Gainscha
//
func (gainscha *Gainscha) GetStatus(deviceID string) *Gainscha {
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
	request := gainscha.request
	request.Method = "POST"
	target := gainscha.Api.List["GetStatus"]
	resp := request.Send(fmt.Sprintf("%v?%v", target, params))
	if nil != resp {
		gainscha.err = resp
	}
	return gainscha
}

//
// DelDev
//	@Description:删除打印机
//	@See:https://dev.poscom.cn/openapi?deldev
//  @receiver gainscha
//  @param deviceID
//  @return *Gainscha
//
func (gainscha *Gainscha) DelDev(deviceID string) *Gainscha {
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
	request := gainscha.request
	request.Method = "POST"
	target := gainscha.Api.List["DelDev"]
	resp := request.Send(fmt.Sprintf("%v?%v", target, params))
	if nil != resp {
		gainscha.err = resp
	}
	return gainscha
}

//
// Group
//  @Description: 获取全部分组
//  @receiver gainscha
//  @return *Gainscha
//
func (gainscha *Gainscha) Group() *Gainscha {
	gainscha.initialize()
	reqTime := src.Time()
	securityCodeParams := src.ArgsType2String(
		gainscha.MemberCode,
		reqTime,
		gainscha.ApiSecretKey,
	)
	securityCode := gainscha.securityCode(securityCodeParams...)
	params := gainscha.generateParam(reqTime, securityCode, nil)
	request := gainscha.request
	request.Method = "GET"
	target := gainscha.Api.List["Group"]
	resp := request.Send(fmt.Sprintf("%v?%v", target, params))
	if nil != resp {
		gainscha.err = resp
	}
	return gainscha
}

//
// ListTemplate
//	@Description:模板列表
//	@See:https://dev.poscom.cn/openapi?templetList
//  @receiver gainscha
//  @return *Gainscha
//
func (gainscha *Gainscha) ListTemplate() *Gainscha {
	gainscha.initialize()
	reqTime := src.Time()
	securityCodeParams := src.ArgsType2String(
		gainscha.MemberCode,
		reqTime,
		gainscha.ApiSecretKey,
	)
	securityCode := gainscha.securityCode(securityCodeParams...)
	params := gainscha.generateParam(reqTime, securityCode, nil)
	request := gainscha.request
	request.Method = "POST"
	target := gainscha.Api.List["ListTemplate"]
	resp := request.Send(fmt.Sprintf("%v?%v", target, params))
	if nil != resp {
		gainscha.err = resp
	}
	return gainscha
}

//
// AddGroup
//  @Description: 添加（打印机）分组
//  @receiver gainscha
//  @param grpName
//  @return *Gainscha
//
func (gainscha *Gainscha) AddGroup(grpName string) *Gainscha {
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
	request := gainscha.request
	request.Method = "POST"
	target := gainscha.Api.List["AddGroup"]
	resp := request.Send(fmt.Sprintf("%v?%v", target, params))
	if nil != resp {
		gainscha.err = resp
	}
	return gainscha
}

//
// ListException
//  @Description: 查询打印异常信息
//  @receiver gainscha
//  @param deviceNo
//  @param start
//  @param end
//  @return *Gainscha
//
func (gainscha *Gainscha) ListException(deviceNo string, start string, end string) *Gainscha {
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
	request := gainscha.request
	request.Method = "GET"
	target := gainscha.Api.List["ListException"]
	resp := request.Send(fmt.Sprintf("%v?%v", target, params))
	if nil != resp {
		gainscha.err = resp
	}
	return gainscha
}

//
// EditGroup
//  @Description: 修改（打印机）分组名称
//  @receiver gainscha
//  @param grpId
//  @param grpName
//  @return *Gainscha
//
func (gainscha *Gainscha) EditGroup(grpId string, grpName string) *Gainscha {
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
	request := gainscha.request
	request.Method = "POST"
	target := gainscha.Api.List["EditGroup"]
	resp := request.Send(fmt.Sprintf("%v?%v", target, params))
	if nil != resp {
		gainscha.err = resp
	}
	return gainscha
}

//
// DelGroup
//  @Description: 删除（打印机）分组名称
//  @receiver gainscha
//  @param grpId
//  @return *Gainscha
//
func (gainscha *Gainscha) DelGroup(grpId string) *Gainscha {
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
	request := gainscha.request
	request.Method = "POST"
	target := gainscha.Api.List["DelGroup"]
	resp := request.Send(fmt.Sprintf("%v?%v", target, params))
	if nil != resp {
		gainscha.err = resp
	}
	return gainscha
}

//
// QueryState
//	@Description:查询打印任务状态
//	@See:https://dev.poscom.cn/openapi?queryState
//  @receiver gainscha
//  @param msgNo
//  @return *Gainscha
//
func (gainscha *Gainscha) QueryState(msgNo string) *Gainscha {
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
	request := gainscha.request
	request.Method = "GET"
	target := gainscha.Api.List["QueryState"]
	resp := request.Send(fmt.Sprintf("%v?%v", target, params))
	if nil != resp {
		gainscha.err = resp
	}
	return gainscha
}

//
// DeleteLogo
//	@Description:删除NVLogo（票据机）
//	@See:https://dev.poscom.cn/openapi?deleteLogo
//  @receiver gainscha
//  @param deviceID
//  @return *Gainscha
//
func (gainscha *Gainscha) DeleteLogo(deviceID string) *Gainscha {
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
	request := gainscha.request
	request.Method = "POST"
	target := gainscha.Api.List["DeleteLogo"]
	resp := request.Send(fmt.Sprintf("%v?%v", target, params))
	if nil != resp {
		gainscha.err = resp
	}
	return gainscha
}

//
// SetPushUrl
//  @Description: 设置接受服务器信息推送地址
//  @receiver gainscha
//  @param pushUrl
//  @return *Gainscha
//
func (gainscha *Gainscha) SetPushUrl(pushUrl string) *Gainscha {
	gainscha.initialize()
	reqTime := src.Time()
	securityCodeParams := src.ArgsType2String(
		gainscha.MemberCode,
		reqTime,
		gainscha.ApiSecretKey,
		pushUrl,
	)
	securityCode := gainscha.securityCode(securityCodeParams...)
	otherParams := map[string]string{
		"pushUrl": pushUrl,
	}
	params := gainscha.generateParam(reqTime, securityCode, otherParams)
	request := gainscha.request
	request.Method = "POST"
	target := gainscha.Api.List["SetPushUrl"]
	resp := request.Send(fmt.Sprintf("%v?%v", target, params))
	if nil != resp {
		gainscha.err = resp
	}
	return gainscha
}

//
// SetLogo
//  @Description: 设置 logo
//  @receiver gainscha
//  @param deviceID
//  @param imgUrl
//  @return *Gainscha
//
func (gainscha *Gainscha) SetLogo(deviceID string, imgUrl string) *Gainscha {
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
		"imgUrl":   imgUrl,
	}
	params := gainscha.generateParam(reqTime, securityCode, otherParams)
	request := gainscha.request
	request.Method = "POST"
	target := gainscha.Api.List["SetLogo"]
	resp := request.Send(fmt.Sprintf("%v?%v", target, params))
	if nil != resp {
		gainscha.err = resp
	}
	return gainscha
}

//func (gainscha *Gainscha) CancelPrint(deviceID string)
