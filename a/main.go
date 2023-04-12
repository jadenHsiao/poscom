package main

import (
	"fmt"
	"github.com/jadenHsiao/poscom"
)

func main() {
	poscom := poscom.NewGainscha("GFTF57YONDFPJU2G1DV004YZTMD9LJAO", "FC2A29F38B18C33C745504C3F6F668AD")
	fmt.Println(poscom.Device().DeviceList())
	fmt.Println(poscom.Setting().SetPushUrl("1111"))
}
