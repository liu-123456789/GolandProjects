package main

import "fmt"

// 自定义人类结构
type Preson struct {
	name  string
	age   uint
	sex   string
	fight uint
}

func main() {
	p1 := Preson{
		name:  "zhangsan",
		age:   18,
		sex:   "男",
		fight: 10,
	}
	fmt.Printf("%v\n", p1)
	p2 := Preson{
		"lisi",
		19,
		"男",
		22,
	}
	fmt.Printf("%v\n", p2)
	p3 := struct {
		name  string
		age   uint
		sex   string
		fight uint
	}{"wangwu", 19, "nan", 20}
	fmt.Println(p3)
	p2 = p3
	fmt.Println("p2", p2)
}
