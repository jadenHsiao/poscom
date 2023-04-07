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
	"net/url"

	"github.com/jadenHsiao/poscom/kernels"
	"github.com/jadenHsiao/poscom/utils"
)

// Group
// @Description:打印机分组结构体
type Group struct {
	list        *List
	deleteGroup *DeleteGroup
	editGroup   *EditGroup
	addGroup    *AddGroup
	kernels.Base
}

// NewGroup
//
//	@Description: 实例化打印机分组机构体
//	@param config
//	@return *Group
func NewGroup(config *kernels.Config) *Group {
	group := new(Group)
	group.Base = *kernels.NewBase(config)
	return group
}

// GroupList
//
//	@Description:查询打印机分组列表
//	@receiver group
//	@return *List
//	@return error
func (group *Group) GroupList() (*List, error) {
	securityCodeParams := utils.ToStrArr(
		group.Config.MemberCode,
		group.Timestamp,
		group.Config.ApiKey,
	)
	securityCode := utils.SecurityCode(securityCodeParams...)
	params := map[string]string{
		"reqTime":      group.Timestamp,
		"memberCode":   group.Base.Config.MemberCode,
		"securityCode": securityCode,
	}
	result, err := NewList().Exec(
		utils.GenerateParam(utils.ParseParam(params), nil),
	)
	return result, err
}

// DeleteGroup
//
//	@Description:删除打印机分组
//	@receiver group
//	@param grpID
//	@return *DeleteGroup
//	@return error
func (group *Group) DeleteGroup(grpID string) (*DeleteGroup, error) {
	securityCodeParams := utils.ToStrArr(
		group.Config.MemberCode,
		group.Timestamp,
		group.Config.ApiKey,
		grpID,
	)
	securityCode := utils.SecurityCode(securityCodeParams...)
	params := map[string]string{
		"reqTime":      group.Timestamp,
		"memberCode":   group.Base.Config.MemberCode,
		"securityCode": securityCode,
		"grpID":        grpID,
	}
	result, err := NewDeleteGroup().Exec(
		utils.GenerateParam(utils.ParseParam(params), nil),
	)
	return result, err
}

// EditGroup
//
//	@Description: 修改打印机分组名称
//	@receiver group
//	@param grpID
//	@param grpName
//	@return *EditGroup
//	@return error
func (group *Group) EditGroup(grpID string, grpName string) (*EditGroup, error) {
	securityCodeParams := utils.ToStrArr(
		group.Config.MemberCode,
		group.Timestamp,
		group.Config.ApiKey,
		grpID,
	)
	securityCode := utils.SecurityCode(securityCodeParams...)
	params := map[string]string{
		"reqTime":      group.Timestamp,
		"memberCode":   group.Base.Config.MemberCode,
		"securityCode": securityCode,
		"grpID":        grpID,
		"grpName":      url.QueryEscape(grpName),
	}
	result, err := NewEditGroup().Exec(
		utils.GenerateParam(utils.ParseParam(params), nil),
	)
	return result, err
}

//
// AddGroup
//  @Description: 新增打印机分组
//  @receiver group
//  @param grpName
//  @return *AddGroup
//  @return error
//
func (group *Group) AddGroup(grpName string) (*AddGroup, error) {
	securityCodeParams := utils.ToStrArr(
		group.Config.MemberCode,
		group.Timestamp,
		group.Config.ApiKey,
	)
	securityCode := utils.SecurityCode(securityCodeParams...)
	params := map[string]string{
		"reqTime":      group.Timestamp,
		"memberCode":   group.Base.Config.MemberCode,
		"securityCode": securityCode,
		"grpName":      url.QueryEscape(grpName),
	}
	result, err := NewAddGroup().Exec(
		utils.GenerateParam(utils.ParseParam(params), nil),
	)
	return result, err
}
