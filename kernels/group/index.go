package group

import (
	"github.com/jadenHsiao/poscom/kernels"
	"github.com/jadenHsiao/poscom/utils"
)

type Group struct {
	list *List
	kernels.Base
}

func NewGroup(config *kernels.Config) *Group {
	group := new(Group)
	group.Base = *kernels.NewBase(config)
	return group
}

func (group *Group) GetGroupList() {
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
	NewList().Group(
		utils.GenerateParam(utils.ParseParam(params), nil),
	)

}
