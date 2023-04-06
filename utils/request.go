// Package utils
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
package utils

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

//
// Send
//  @Description: 发起网络请求
//  @param method string 请求类型
//  @param uri string 目标 url
//  @return result
//  @return err
//
func Send(method string, uri string) (result []byte, err error) {
	var resp *http.Response
	if "GET" == method {
		resp, err = http.Get(uri)
	} else if "POST" == method {
		resp, err = http.Post(uri, "application/x-www-form-urlencoded", nil)
	}
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

//
// GenerateParam
//  @Description: 生成请求参数
//  @param paramList
//  @param otherParams
//  @return string
//
func GenerateParam(paramList []string, otherParams map[string]string) string {
	for key, val := range otherParams {
		item := fmt.Sprintf("%v=%v", key, val)
		paramList = append(paramList, item)
	}
	return strings.Join(paramList, "&")
}

//
// ParseParam
//  @Description: 生成 uri 参数
//  @param params
//  @return []string
//
func ParseParam(params map[string]string) []string {
	var result []string
	for key, val := range params {
		result = append(result, fmt.Sprintf("%v=%v", key, val))
	}
	return result
}

//
// Submit
//  @Description: 发起请求并解析到数据
//  @param uri
//  @param params
//  @param method
//  @return []byte
//  @return error
//
func Submit(uri string, params string, method string) ([]byte, error) {
	uri = fmt.Sprintf("%v?%v", uri, params)
	ctx, err := Send(method, uri)
	if nil != err {
		return nil, err
	}
	return ctx, nil
}
