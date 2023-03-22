// Package src
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
package src

const (
	//  接口请求根域名
	ApiHost = "http://api.poscom.cn"
	//  请求端口
	PORT = "80"
	//  请求路径
	BasePath = "/apisc"
)

//
//  Api
//  @Description: `Api`结构体
//
type Api struct {
	List map[string]string
}

//
// Initialize
//  @Description: 初始化`Api`结构体
//  @receiver api
//  @return *Api
//
func (api *Api) Initialize() *Api {
	var list map[string]string
	list = map[string]string{
		"SendMsg":       "sendMsg",
		"QueryState":    "queryState",
		"GetStatus":     "getStatus",
		"ListException": "listException",
		"AddGroup":      "addgroup",
		"Group":         "group",
		"EditGroup":     "editgroup",
		"DelGroup":      "delgroup",
		"AddDev":        "adddev",
		"Device":        "device",
		"ListDevice":    "listDevice",
		"EditDev":       "editdev",
		"DelDev":        "deldev",
		"CancelPrint":   "cancelPrint",
		"SetLogo":       "setLogo",
		"DeleteLogo":    "deleteLogo",
	}
	for key, val := range list {
		list[key] = ApiHost + ":" + PORT + BasePath + "/" + val
	}
	api.List = list
	return api
}
