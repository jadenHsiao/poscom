# Poscom SDK for Golang

![version](https://img.shields.io/badge/version-v1-green)
![license](https://img.shields.io/badge/license-MulanPSL--2.0-blue)

这是一个使用Golang开发的佳博打印机SDK。
> 本SDK是基于佳博云打印Web API文档（v1.9.7）和[在线文档](https://dev.poscom.cn/docs)所开发

## 快速开始

1. 安装

````go
import "github.com/jadenHsiao/poscom"
````
2. 开始使用（以下是一个使用本SDK获取设备列表的例子）：
````go
// 传入开发者的`Api`编码和`Api`秘钥
gainscha := poscom.NewGainscha(
	"GFTF57**********4YZTMD9LJAO", 
	"FC2A29**********4C3F6F668AD")
// 调用设备模块下的设备列表
gainscha.Device().DeviceList()
````

## 模块说明

| 模块名称          | 名称      | 简介                    |
|---------------|---------|-----------------------|
| Device        | 设备模块    | 设备的增、删、改、查            |
| DeviceSetting | 打印机设置模块 | 设置音量、切换播报类型、修改网络类型    |
| Group         | 打印机分组模块 | 打印机分组的增、删、改、查         |
| Logo          | Logo模块  | 打印Logo的设置和删除          |
| Printer       | 打印机模块   | 发送、查询、取消打印任务，获取打印错误信息 |
| Setting       | 其他设置模块  | 设置回调和升级硬件             |
| Template      | 模板模块    | 模板列表                  |

### 设备模块
| 方法           | 简介   | 
|------------|-----|
| DeviceList |  设备列表 |
| AddDevice  |  添加设备 |                                               
| DeleteDevice  |  删除设备 | 
| EditDevice  | 编辑设备 |
| DeviceDetail  |  设备详情 |

### 打印机设置模块
| 方法         |  简介   |
|------------|------|
| SetNet |  设置网络链接方式 |
| SendVolume  |  打印机音量设置 |
| SendVoiceType  |  打印机切换播报类型 |

### 打印机分组模块
| 方法           | 简介   | 
|------------|-----|
| GroupList |  查询打印机分组列表 |
| DeleteGroup  |  删除打印机分组 |                                               
| EditGroup  |  修改打印机分组名称 | 
| AddGroup  | 新增打印机分组 |

### Logo模块
| 方法         |  简介   |
|------------|------|
| SetLogo |  设置logo |
| DeleteLogo  |  删除打印机NVLogo|

### 打印机模块
| 方法           | 简介       | 
|------------|----------|
| QueryState | 查询打印任务状态 |
| ListException  | 查询打印异常信息 |                                               
| GetStatus  | 获取打印机状态  | 
| CancelPrint  | 取消打印     |
| SendMsg  | 发送数据到打印机     |

### 打印机设置模块
| 方法         |  简介   |
|------------|------|
| SetPushUrl |  设置接受服务器信息推送地址 |
| RomUpdate  |  票据打印机固件升级 |

### 模板模块
| 方法         |  简介   |
|------------|------|
| ListTemplate |  模板列表 |


## 贡献
* 在API列表中查看哪些API未实现
* 提交issue，描述需要贡献的内容
* 完成更改后，提交PR

## 联系我
* 微信公众号：CodesHub码坞
* 技术blog：www.jytype.cn
* facebook：Jason Shaw

## License
木兰宽松许可证， 第2版
