package main

import "fmt"

func main() {

	fmt.Printf("你喜欢什么水果")
	var fruit string
	//接受标准输入，获取一个水果名称
	fmt.Scanf("%s", &fruit)
	//fmt.Println(fruit)
	switch fruit {
	case "apple":
		fmt.Println("i like apple")
		fallthrough //继续执行下面得分支
	case "banana":
		fmt.Println("i like banana")
	case "pear":
		fmt.Println("i like pear")
	case "mango":
		fmt.Println("i like mango")
	default:
		fmt.Println("are you kidding me?")
	}
}
