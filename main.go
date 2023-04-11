//package poscom
//
//import (
//	poscom "github.com/jadenHsiao/poscom/src/kernels"
//)
//
//func main() {
//
//	gainscha := &poscom.Gainscha{
//		ApiSecretKey: "GFTF57YONDFPJU2G1DV004YZTMD9LJAO",
//		MemberCode:   "FC2A29F38B18C33C745504C3F6F668AD",
//	}
//	//fmt.Println(gainscha.SendVolume("00135465930408666", 0))
//	//fmt.Println(gainscha.ListDevice())
//	//fmt.Println(gainscha.Device("00135465930408666"))
//	//fmt.Println(gainscha.DelDev("00135465930408666"))
//	//fmt.Println(gainscha.GetStatus("00135465930408666"))
//	//fmt.Println(gainscha.Group())
//	//fmt.Println(gainscha.ListTemplate())
//	//fmt.Println(gainscha.ListException("00135465930408666", "", ""))
//	//fmt.Println(gainscha.AddGroup("分组2"))
//	//fmt.Println(gainscha.EditGroup("46003", "分组3"))
//	//fmt.Println(gainscha.DelGroup("46003"))
//	//fmt.Println(gainscha.QueryState("1000001225881146"))
//	//fmt.Println(gainscha.DeleteLogo("00135465930408666"))
//	//fmt.Println(gainscha.SetPushUrl("00135465930408666"))
//	//fmt.Println(gainscha.SetLogo("00135465930408666", "1111"))
//}
package main

import (
	"fmt"
	"github.com/jadenHsiao/poscom/kernels"
	"github.com/jadenHsiao/poscom/kernels/setting"
)

func main() {
	cfg := kernels.Config{
		ApiKey:     "GFTF57YONDFPJU2G1DV004YZTMD9LJAO",
		MemberCode: "FC2A29F38B18C33C745504C3F6F668AD",
	}
	// 打印机分组
	//fmt.Println(group.NewGroup(&cfg).GroupList())
	//fmt.Println(group.NewGroup(&cfg).DeleteGroup("46386"))
	//fmt.Println(group.NewGroup(&cfg).EditGroup("46005", "分组2"))
	//fmt.Println(group.NewGroup(&cfg).AddGroup("分组2"))
	// logo
	//fmt.Println(logo.NewLogo(&cfg).SetLogo("00135465930408666", "1111"))
	//fmt.Println(logo.NewLogo(&cfg).DeleteLogo("00135465930408666"))
	// 网络设置
	//fmt.Println(devicesetting.NewDeviceSetting(&cfg).SendVolume("00135465930408666", 0))
	//fmt.Println(devicesetting.NewDeviceSetting(&cfg).SetNet("00135465930408666", "wifi"))
	//fmt.Println(devicesetting.NewDeviceSetting(&cfg).SendVoiceType("00135465930408666", 0))
	// 打印机相关
	//fmt.Println(device.NewDevice(&cfg).DeviceList())
	//fmt.Println(device.NewDevice(&cfg).DeviceDetail("00135465930408666"))
	//fmt.Println(device.NewDevice(&cfg).DeleteDevice("00135465930408666"))

	// 设置相关
	fmt.Println(setting.NewSetting(&cfg).SetPushUrl("http://www.baidu.co"))
}
