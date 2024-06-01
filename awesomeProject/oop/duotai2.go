package main

import "fmt"

type Presonc struct {
	name  string
	age   uint
	sex   string
	fight uint
}

type HUman interface {
	getName() string
	setage(age uint)
}

func dosths(h HUman) {
	h.getName()
	h.setage(101)
	fmt.Println(h)
}

// 为Person提供方法
func (p Presonc) getName() string {
	return p.name
}

// 若要修改结构内容使用”*“指针，不想修改结构内容使用对象
func (this *Presonc) setage(age uint) {
	this.age = age
}

func main() {
	p1 := Presonc{
		name:  "zhangsan",
		age:   18,
		sex:   "男",
		fight: 10,
	}
	dosths(&p1)
	//fmt.Println(p1.getName())
	//p1.setage(19)
	//fmt.Println(p1)
}
