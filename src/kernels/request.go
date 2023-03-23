package kernels

import (
	"net/http"
)

type Request struct {
	Method      string
	Data        map[string]interface{}
	Header      map[string]interface{}
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

func (req *Request) send(url string, data map[string]interface{}) {
	req.initialize()
	if false == req.checkMethod() {

	}
	if "GET" == req.Method {

	} else if "POST" == req.Method {

	}

	http.NewRequest(req.Method, url, nil)
}
