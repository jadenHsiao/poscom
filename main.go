package main

import (
	"fmt"
	"github.com/jadenHsiao/poscom/src/kernels"
)

func main() {

	gainscha := kernels.Gainscha{
		ApiSecretKey: "GFTF57YONDFPJU2G1DV004YZTMD9LJAO",
		MemberCode:   "FC2A29F38B18C33C745504C3F6F668AD",
	}
	//gainscha.ListDevice()
	//fmt.Println(gainscha.Device("00135465930408666"))
	//fmt.Println(gainscha.Device("00135465930408666"))
	fmt.Println(gainscha.AddGroup("分组2"))
}
