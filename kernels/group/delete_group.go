package group

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
	"encoding/json"
	"fmt"
	"github.com/jadenHsiao/poscom/utils"
)

const DeleteGroupUri = "http://api.poscom.cn:80/apisc/delgroup"

//
//  DelGroup
//  @Description:删除打印机分组数据返回结构体
//
type DeleteGroup struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

//
// NewDeleteGroup
//  @Description: 实例化删除打印机分组
//  @return *DeleteGroup
//
func NewDeleteGroup() *DeleteGroup {
	deleteGroup := new(DeleteGroup)
	return deleteGroup
}

//
// Submit
//  @Description: 删除账号下打印机分组
//  @receiver deleteGroup
//  @param params
//  @return *DeleteGroup
//  @return error
//
func (deleteGroup *DeleteGroup) Submit(params string) (*DeleteGroup, error) {
	uri := fmt.Sprintf("%v?%v", DeleteGroupUri, params)
	ctx, err := utils.Send("POST", uri)
	if nil != err {
		return nil, err
	}
	var result *DeleteGroup
	err = json.Unmarshal(ctx, &result)
	if nil != err {
		return nil, err
	}
	return result, nil
}
