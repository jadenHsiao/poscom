package main

import (
	"fmt"
	"github.com/jadenHsiao/poscom"
)

func main() {
	gainscha := poscom.NewGainscha("GFTF57YONDFPJU2G1DV004YZTMD9LJAO", "FC2A29F38B18C33C745504C3F6F668AD")
	//fmt.Println(gainscha.Device().DeviceList())
	//fmt.Println(gainscha.Setting().SetPushUrl("1111"))
	//fmt.Println(gainscha.Printer().QueryState("1000001896147286"))
	//fmt.Println(gainscha.Printer().ListException("00135465930408666", "", ""))
	//fmt.Println(gainscha.Printer().GetStatus(""))
	//fmt.Println(gainscha.Printer().GetStatus(""))
	//fmt.Println(gainscha.Printer().CancelPrint("00135465930408666", "0"))
	//fmt.Println(gainscha.Device().AddDevice("1111111", "231232", make(map[string]string)))
	//fmt.Println(gainscha.Device().EditDevice("00135465930408666", map[string]string{"devName": "打印机"}))
	fmt.Println(gainscha.Template().ListTemplate())
}
