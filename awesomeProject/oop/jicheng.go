package main

import "fmt"

type Persone struct {
	name  string
	age   uint
	sex   string
	fight uint
}

func (p Persone) getName() string {
	return p.name
}

func (this *Persone) setage(age uint) {
	this.age = age
}

// 继承Persone
type SuperMan struct {
	Persone
	cnl  string
	name string
}

// 非内嵌Persone
type SuperMan2 struct {
	p    Persone //非内嵌，属性定义
	cul  string
	name string
}

func main() {
	s1 := SuperMan{
		Persone{
			"zhangsan",
			30,
			"man",
			1000,
		},
		"大力士",
		"王五",
	}
	fmt.Printf("%v\n", s1)
	s1.setage(100)
	fmt.Println(s1)
	fmt.Println(s1.name, s1.getName(), s1.age)
	s2 := SuperMan2{
		p: Persone{
			"lisi",
			20,
			"man",
			10,
		},
		cul:  "历史",
		name: "wangwu",
	}
	fmt.Println(s2.p.getName())
}
