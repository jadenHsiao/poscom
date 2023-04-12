package setting

import "github.com/jadenHsiao/poscom/utils"

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

const RomUpdateUri = "http://api.poscom.cn/apisc/romupdate"

//
//  RomUpdate
//  @Description:票据打印机固件升级数据返回结构体
//
type RomUpdate struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

//
// NewRomUpdate
//  @Description: 实例化票据打印机固件升级
//  @return *RomUpdate
//
func NewRomUpdate() *RomUpdate {
	romUpdate := new(RomUpdate)
	return romUpdate
}

//
// Exec
//  @Description: 发起请求
//  @receiver romUpdate
//  @param params
//  @return *RomUpdate
//  @return error
//
func (romUpdate *RomUpdate) Exec(params string) (*RomUpdate, error) {
	ctx, err := utils.Send(RomUpdateUri, params, "POST")
	if err != nil {
		return nil, err
	}
	result, err := utils.JsonDecode(ctx, new(RomUpdate))
	return result.(*RomUpdate), err
}
