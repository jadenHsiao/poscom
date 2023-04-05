package api

import (
	"fmt"
	"reflect"
)

package main

import (
"fmt"
"reflect"
)
//定义Monster结构体
type Monster struct {
	Name string `json:"name"`
	Age int `json:"monster_age"`
	Score float32 `json:"成绩"`
	Sex string
}
//方法，返回两个数的和
func (s Monster) GetSum(n1, n2 int) int {
	return n1 + n2
}
//方法，接受4个值，给s赋值
func (s Monster) Set(name string,age int,score float32,sex string) {
	s.Name = name
	s.Age = age
	s.Score = score
	s.Sex = sex
}
//方法，显示s的值
func (s Monster) Print() {
	fmt.Println("----start----")
	fmt.Println(s)
	fmt.Println("----end----")
}
func TestStruct(a interface{}) {

	typ := reflect.TypeOf(a)
	val := reflect.ValueOf(a) //获取reflect.Type类型

	kd := val.Kind() //获取到a对应的类别
	if kd != reflect.Struct {
		fmt.Println("expect struct")
		return
	}
	//获取到该结构体有几个字段
	num := val.NumField()
	fmt.Printf("该结构体有%d个字段\n", num) //4个

	//遍历结构体的所有字段
	for i := 0; i < num; i++ {
		fmt.Printf("Field %d:值=%v\n",i,val.Field(i))
		//获取到struct标签，需要通过reflect.Type来获取tag标签的值
		tagVal := typ.Field(i).Tag.Get("json")
		//如果该字段有tag标签就显示，否则就不显示
		if tagVal != ""{
			fmt.Printf("Field %d:tag=%v\n",i,tagVal)
		}
	}
	//获取到该结构体有多少个方法
	numOfMethod := val.NumMethod()
	fmt.Printf("struct has %d methods\n",numOfMethod)

	//方法的排序默认是按照函数名的顺序（ASCII码）
	val.Method(1).Call(nil) //获取到第二个方法，调用它

	//调用结构体的第1个方法，Method(0)
	var params []reflect.Value
	params = append(params,reflect.ValueOf(10))
	params = append(params,reflect.ValueOf(40))
	//传入的参数是 []reflect.Value，返回 []reflect.Value
	res := val.Method(0).Call(params)
	//返回结果，返回的结果是 []reflect.Value
	fmt.Println("返回的结果 res=",res[0].Int())

}

func main() {
	//创建一个Monster实例
	var a Monster = Monster{
		Name:"张三丰",
		Age:99,
		Score:80.5,
	}
	//将Monster实例传递给TestStruct函数
	TestStruct(a)
}
