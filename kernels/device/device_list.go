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
import "github.com/jadenHsiao/poscom/utils"

const DeviceListUri = "http://api.poscom.cn:80/apisc/listDevice"

//
//  List
//  @Description:打印机列表数据返回结构体
//
type List struct {
	Code       int          `json:"code"`
	Msg        string       `json:"msg"`
	DeviceList []DeviceList `json:"deviceList"`
}

//
//  DeviceList
//  @Description:打印机列表数组结构体
//
type DeviceList struct {
	Id     string `json:"id"`
	Imsi   string `json:"imsi"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

//
// NewList
//  @Description: 实例化打印机列表数据返回结构体
//  @return *List
//
func NewList() *List {
	list := new(List)
	return list
}

//
// Exec
//  @Description: 执行请求
//  @receiver list
//  @param params
//  @return *List
//  @return error
//
func (list *List) Exec(params string) (*List, error) {
	ctx, err := utils.Submit(DeviceListUri, params, "POST")
	if err != nil {
		return nil, err
	}
	result, err := utils.JsonDecode(ctx, new(List))
	return result.(*List), err
}
