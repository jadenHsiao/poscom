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

const EditGroupUri = "http://api.poscom.cn:80/apisc/editgroup"

//
//  EditGrp
//  @Description:修改打印机分组名称数据返回结构体
//
type EditGroup struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

//
// NewEditGroup
//  @Description: 实例化修改打印机分组名称
//  @return *EditGroup
//
func NewEditGroup() *EditGroup {
	editGrp := new(EditGroup)
	return editGrp
}

//
// Submit
//  @Description: 发起请求
//  @receiver editGrp
//  @param params
//  @return *EditGrp
//  @return error
//
func (editGroup *EditGroup) Submit(params string) (*EditGroup, error) {
	uri := fmt.Sprintf("%v?%v", EditGroupUri, params)
	ctx, err := utils.Send("POST", uri)
	if nil != err {
		return nil, err
	}
	var result *EditGroup
	err = json.Unmarshal(ctx, &result)
	if nil != err {
		return nil, err
	}
	return result, nil
}
