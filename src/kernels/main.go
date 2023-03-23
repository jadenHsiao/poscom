// Package kernels
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
package kernels

import (
	"fmt"
	"github.com/jadenHsiao/poscom/src"
)

//
//  Gainscha
//  @Description: 佳博打印结构体
//
type Gainscha struct {
	Api          *src.Api
	ApiSecretKey string
	MemberCode   string
}

//
// initialize
//  @Description: 初始化佳博打印结构体
//  @receiver gainscha
//
func (gainscha *Gainscha) initialize() {
	gainscha.Api = new(src.Api)
	gainscha.Api.Initialize()
}

//
// securityCode
//  @Description: 根据官方规则创建秘钥
//  @receiver gainscha
//  @param args
//  @return string
//
func (gainscha *Gainscha) securityCode(args ...string) string {
	var beforeStr string
	for _, val := range args {
		beforeStr += val
	}
	return src.Md5(beforeStr)
}

func (gainscha *Gainscha) A() *src.Api {
	gainscha.initialize()
	fmt.Println(gainscha.Api.List["Group"])
	return gainscha.Api
}
