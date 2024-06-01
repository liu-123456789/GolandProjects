package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x1 int //int整型类型，uint只能是正整数
	fmt.Println(x1)
	var x2 float32 // go语言float包括这两种类型，要不然就，float32 或者使用float64区别是精度不一样
	fmt.Println(reflect.TypeOf(x2))
	var x3 = "hello" //字符串类型
	fmt.Println(x3)
	var x4, x5 = 100, true //可以赋值多个变量
	fmt.Println(reflect.TypeOf(x4), reflect.TypeOf(x5))
	var x6 complex64 = 3 + 2i
	fmt.Println(x6)
	x7, x8 := 123, "holle"
	fmt.Println(x7, x8)

}
