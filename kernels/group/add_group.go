package group

import (
	"encoding/json"
	"fmt"
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

const AddGroupUri = "http://api.poscom.cn:80/apisc/addgroup"

//
//  AddGroup
//  @Description:添加打印机分组数据返回结构体
//
type AddGroup struct {
	Code  int    `json:"code"`
	Msg   string `json:"msg"`
	GrpID int    `json:"grpID"`
}

//
// NewAddGroup
//  @Description: 实例化添加打印机分组
//  @return *AddGroup
//
func NewAddGroup() *AddGroup {
	addGroup := new(AddGroup)
	return addGroup
}

//
// Submit
//  @Description: 添加打印机分组
//  @receiver addGroup
//  @param params
//  @return *AddGroup
//  @return error
//
func (addGroup *AddGroup) Submit(params string) (*AddGroup, error) {
	uri := fmt.Sprintf("%v?%v", AddGroupUri, params)
	ctx, err := utils.Send("POST", uri)
	if nil != err {
		return nil, err
	}
	var result *AddGroup
	err = json.Unmarshal(ctx, &result)
	if nil != err {
		return nil, err
	}
	return result, nil
}
