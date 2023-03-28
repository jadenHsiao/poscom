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
	"io"
	"net/http"

	"github.com/jadenHsiao/poscom/src"
)

// Request
// @Description:请求结构体
type Request struct {
	Method   string
	Response []byte
}

// Send
//
//	@Description: 发送请求
//	@receiver req
//	@param url
//	@return result
//	@return errContent
func (req *Request) Send(url string) (result *src.PoscomError) {
	var resp *http.Response
	var err error
	if "GET" == req.Method {
		resp, err = http.Get(url)
	} else if "POST" == req.Method {
		resp, err = http.Post(url, "application/x-www-form-urlencoded", nil)
	}
	if err != nil {
		return src.NewError("101001", "")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return src.NewError("101002", err.Error())
	}
	req.Response = body
	return nil

}
