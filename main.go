package main

import (
	"fmt"

	"github.com/jadenHsiao/poscom/src/kernels"
)

func main() {

	gainscha := &kernels.Gainscha{
		ApiSecretKey: "GFTF57YONDFPJU2G1DV004YZTMD9LJAO",
		MemberCode:   "FC2A29F38B18C33C745504C3F6F668AD",
	}
	fmt.Println(gainscha.ListDevice().ToMap())

	//fmt.Println(gainscha.ListDevice())
	//fmt.Println(gainscha.Device("00135465930408666"))
	//fmt.Println(gainscha.DelDev("00135465930408666"))
	//fmt.Println(gainscha.GetStatus("00135465930408666"))
	//fmt.Println(gainscha.Group())
	//fmt.Println(gainscha.ListTemplate())
	//fmt.Println(gainscha.ListException("00135465930408666", "", ""))
	//fmt.Println(gainscha.AddGroup("分组2"))
	//fmt.Println(gainscha.EditGroup("46003", "分组3"))
	//fmt.Println(gainscha.DelGroup("46003"))
	//fmt.Println(gainscha.QueryState("1000001225881146"))
	//fmt.Println(gainscha.DeleteLogo("00135465930408666"))
	//fmt.Println(gainscha.SetPushUrl("00135465930408666"))
	//fmt.Println(gainscha.SetLogo("00135465930408666", "1111"))
}
