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

// @See:https://dev.poscom.cn/debug?editgroup
const EditGroupUri = "http://api.poscom.cn:80/apisc/editgroup"

//
//  EditGrp
//  @Description:修改打印机分组名称结构体
//
type EditGrp struct {
	Code int
	Msg  string
}

//
// NewEditGrp
//  @Description: 实例化修改打印机名称
//  @return *EditGroup
//
func NewEditGrp() *EditGrp {
	editGrp := new(EditGrp)
	return editGrp
}

//
// EditGroup
//  @Description:修改打印机分组名称
//  @receiver editGrp
//  @param params
//  @return *EditGrp
//  @return error
//
func (editGrp *EditGrp) EditGroup(params string) (*EditGrp, error) {
	uri := fmt.Sprintf("%v?%v", EditGroupUri, params)
	ctx, err := utils.Send("POST", uri)
	if nil != err {
		return nil, err
	}
	var result *EditGrp
	err = json.Unmarshal(ctx, &result)
	if nil != err {
		return nil, err
	}
	return result, nil
}
