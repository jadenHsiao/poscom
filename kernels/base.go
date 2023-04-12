package kernels

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

//
// NewBase
//  @Description: 实例化基础结构体
//  @param config
//  @return *Base
//
func NewBase(config *Config) *Base {
	base := new(Base)
	base.Config = *config
	base.Timestamp = strconv.FormatInt(utils.CurrentTimestamp(), 10)
	return base
}
