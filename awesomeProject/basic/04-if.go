package main

import "fmt"

func main() {
	//条件判断结果必须为bool值
	a := 10
	if a < 10 {
		fmt.Println("a小于10")
	} else if a > 10 {
		fmt.Printf("a大于10")
	} else {
		fmt.Println("a等于10")
	}
}
