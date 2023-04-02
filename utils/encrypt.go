// Package utils
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
package utils

import (
	"crypto/md5"
	"fmt"
	"io"
)

//
// Md5
//  @Description: 计算字符串的`md5`值
//  @param str string 加密前的字符串
//  @return string
//
func Md5(str string) string {
	h := md5.New()
	io.WriteString(h, str)
	return fmt.Sprintf("%x", h.Sum(nil))
}

//
// SecurityCode
//  @Description: 创建安全检验码
//  @param args map[]string 集合
//  @return string
//
func SecurityCode(args ...string) string {
	var beforeStr string
	for _, val := range args {
		beforeStr += val
	}
	return Md5(beforeStr)
}
