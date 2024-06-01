package main

import "fmt"

// const 常量 全局定义，不可以修改
const PI float64 = 3.1415926

// 常用来定义枚举，所谓枚举就是固定的整数值
const (
	APPLE  = 0
	BANANA = 1
	PEAR   = 2
	MANGO  = 3
)

const (
	APPLE1 = iota
	BANANA2
	PEAR3
	MANGO4
)

const (
	APPLE11, APPLE12 = iota, iota + 1
	BANANA21, BANANA22
	PEAR31, PEAR33 = iota * 3, iota + 1
	MANGO41, MANGO42
)

func main() {
	fmt.Println(PI)
	fmt.Println(APPLE)
	fmt.Println(BANANA)
	fmt.Println(PEAR)
	fmt.Println(MANGO)
	fmt.Println("1", APPLE1, BANANA2, PEAR3, MANGO4)
	fmt.Println("2=", APPLE11, APPLE12, BANANA21, BANANA22, PEAR31, PEAR33, MANGO41, MANGO42)
}
