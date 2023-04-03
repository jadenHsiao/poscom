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

// @See:https://dev.poscom.cn/debug?group
const GroupListUri = "http://api.poscom.cn:80/apisc/group"

//
//  List
//  @Description:打印机分组列表完整结构体
//
type List struct {
	Code    int       `json:"code"`
	Msg     string    `json:"msg"`
	GrpList []GrpList `json:"grpList"`
}

//
//  GrpList
//  @Description: 打印机分组列表数组结构体
//
type GrpList struct {
	GrpID   int    `json:"grpID"`
	GrpName string `json:"grpName"`
}

//
// NewList
//  @Description: 实例化打印机分组结构体
//  @return *List
//
func NewList() *List {
	list := new(List)
	return list
}

//
// Group
//  @Description: 查询打印机分组列表
//  @receiver list
//  @param params
//  @return *List
//  @return error
//
func (list *List) Group(params string) (*List, error) {
	uri := fmt.Sprintf("%v?%v", GroupListUri, params)
	ctx, err := utils.Send("GET", uri)
	if nil != err {
		return nil, err
	}
	var result *List
	err = json.Unmarshal(ctx, &result)
	if nil != err {
		return nil, err
	}
	return result, nil
}
