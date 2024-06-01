package main

import "fmt"

type Presons struct {
	name  string
	age   uint
	sex   string
	fight uint
}

// 为Person提供方法
func (p Presons) getName() string {
	return p.name
}

// 若要修改结构内容使用”*“指针，不想修改结构内容使用对象
func (this *Presons) setage(age uint) {
	this.age = age
}

func main() {
	p1 := Presons{
		name:  "zhangsan",
		age:   18,
		sex:   "男",
		fight: 10,
	}
	fmt.Println(p1.getName())
	p1.setage(19)
	fmt.Println(p1)
}
