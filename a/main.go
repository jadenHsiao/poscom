package main

import (
	"fmt"
	"github.com/jadenHsiao/poscom"
)

func main() {
	gainscha := poscom.NewGainscha("GFTF57YONDFPJU2G1DV004YZTMD9LJAO", "FC2A29F38B18C33C745504C3F6F668AD")
	fmt.Println(gainscha.Device().DeviceList())
	//fmt.Println(gainscha.Setting().SetPushUrl("1111"))
	//fmt.Println(gainscha.Printer().QueryState("1000001896147286"))
}
