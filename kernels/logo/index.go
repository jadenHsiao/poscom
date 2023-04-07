package logo

import "github.com/jadenHsiao/poscom/kernels"

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

type Logo struct {
	kernels.Base
}

//
// NewLogo
//  @Description: 实例化`logo`相关结构体
//  @param config
//  @return *Logo
//
func NewLogo(config *kernels.Config) *Logo {
	logo := new(Logo)
	logo.Base = *kernels.NewBase(config)
	return logo
}
