# Poscom SDK for Golang

![version](https://img.shields.io/badge/version-v1-green)
![license](https://img.shields.io/badge/license-MulanPSL--2.0-blue)
[![Open Source Love](https://badges.frapsoft.com/os/v1/open-source.svg?v=103)](https://github.com/ellerbrock/open-source-badges/)

这是一个使用Golang开发的佳博打印机SDK。
> 本SDK是基于佳博云打印Web API文档（v1.9.7）和[在线文档](https://dev.poscom.cn/docs)所开发

## 快速开始

1. 安装

````go
import "github.com/jadenHsiao/poscom/src/kernels"
````
2. 开始使用（以下是一个使用本SDK获取设备列表的例子）：
````go
// 传入开发者的`Api`编码和`Api`秘钥
gainscha := &kernels.Gainscha{
    ApiSecretKey: "GFTF57Y*******4YZTMD9LJAO",
    MemberCode:   "FC2A29F*******4C3F6F668AD",
}
// 调用设备列表最后以字符串输出
gainscha.ListDevice().ToString()
````

## 接口一览


1. 函数和接口对应参照

| 函数  | 接口  | 类型  | 状态  |
|-----|-----|-----|-----|
| SendMsg | sendMsg |  |  |
| QueryState | queryState | 查询打印任务状态 | ✅ |
| GetStatus | getStatus | 获取打印机状态 | ✅ |
| ListException | listException | 查询打印异常信息 | ✅ |
| AddGroup | addgroup | 添加（打印机）分组 | ✅ |
| Group | group | 获取全部分组 | ✅ |
| EditGroup | editgroup | 修改（打印机）分组名称 | ✅ |
| DelGroup | delgroup | 删除（打印机）分组名称 | ✅ |
| AddDev | adddev |  |  |
| Device | device | 查询打印机信息 | ✅ |
| ListDevice | listDevice | 查询打印机列表 | ✅ |
| EditDev | editdev |  |  |
| DelDev | deldev | 删除打印机 | ✅ |
| CancelPrint | cancelPrint | 取消打印订单 | ✅ |
| SetLogo | setLogo | 设置 logo | ✅ |
| DeleteLogo | deleteLogo | 删除NVLogo（票据机） | ✅ |
| ListTemplate | listTemplate | 单元格 | ✅ |
| SetPushUrl | setPushUrl | 设置接受服务器信息推送地址 | ✅ |
| SendVolume | sendVolume | 打印机设置音量 | ✅ |
| SetVoiceType | setvoicetype | 语音播放方式设置 | ✅ |
| SetNet | setNet | 网络链接方式设置 | ✅ |

2. 数据处理函数

| 函数  | 功能  |  
|-----|-----|
| ToMap|返回为结构体|
| ToString|返回为字符串|
| ToByte|返回为Byte|

## 目录文件说明

````go
|-- poscom  
    |-- src                 // SDK 目录
        |-- api.go          // `Api`结构体
        |-- err.go          // 错误处理结构体
        |-- functions.go    // 辅助函数
        |-- kernels         // 核心目录
            |-- main.go     // 核心文件
            |-- request.go  // 网络请求文件
````

## 贡献
* 在API列表中查看哪些API未实现
* 提交issue，描述需要贡献的内容
* 完成更改后，提交PR

## License
木兰宽松许可证， 第2版
