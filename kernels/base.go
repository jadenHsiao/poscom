package kernels

import (
	"github.com/jadenHsiao/poscom/utils"
	"strconv"
)

//
//  Base
//  @Description: 基础结构体
//
type Base struct {
	Config    Config
	Timestamp string
}

func NewBase(config *Config) *Base {
	base := new(Base)
	base.Config = *config
	base.Timestamp = strconv.FormatInt(utils.CurrentTimestamp(), 10)
	return base
}
