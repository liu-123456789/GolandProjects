package main

import "fmt"

/*
函数闭包，需要至少两层函数，或者父子函数
子函数访问父函数的变量
*/

func getNextNumber() func() int {
	num := 0
	return func() int {
		num++ //访问父函数的变量
		return num
	}
}

func main() {
	f1 := getNextNumber()
	fmt.Println(f1())
	fmt.Println(f1())
	fmt.Println(f1())
	fmt.Println(f1())
	fmt.Println(f1())
	fmt.Println("--------------------")
	f2 := getNextNumber()
	fmt.Println(f2())
	fmt.Println(f2())
	fmt.Println(f2())
	fmt.Println(f1())
}
