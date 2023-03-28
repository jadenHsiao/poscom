# Poscom SDK for Golang

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

## 接口说明
| API | 请求类型 |
|-----|------|
| 单元格 | 单元格  |
| 单元格 | 单元格  |

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
