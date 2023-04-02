package group

import (
	"fmt"
	"github.com/jadenHsiao/poscom/utils"
)

const GROUP_LIST_URL = "http://api.poscom.cn:80/apisc/group"

type List struct {
}

func NewList() *List {
	list := new(List)
	return list
}

func (list *List) Group(params string) ([]*List, error) {
	uri := fmt.Sprintf("%v?%v", GROUP_LIST_URL, params)
	utils.Send("GET", uri)
}
