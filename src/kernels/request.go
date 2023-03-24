package kernels

import (
	"fmt"
	"io"
	"net/http"
)

type Request struct {
	Method      string
	validMethod []string
}

func (req *Request) initialize() {
	req.validMethod = []string{
		0: "GET",
		1: "POST",
	}
}

//
// checkMethod
//  @Description: 检查请求类型是否合法
//  @receiver req
//  @return bool
//
func (req *Request) checkMethod() bool {
	for _, val := range req.validMethod {
		if val == req.Method {
			return true
		}
	}
	return false
}

func (req *Request) Send(url string, data map[string]interface{}) {
	req.initialize()
	var resp *http.Response
	var err error
	if false == req.checkMethod() {

	}
	if "GET" == req.Method {
		resp, err = http.Get(url)
	} else if "POST" == req.Method {
		resp, err = http.Post(url, "application/x-www-form-urlencoded", nil)
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
