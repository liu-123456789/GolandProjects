package main

import "fmt"

type Ainimal interface {
	sleeping()
	eating()
}

type Cat struct {
	color string
}

type Dog struct {
	color string
}

func (c Cat) sleeping() {
	fmt.Printf("%s cat is sleeping\n", c.color)
}

func (c Cat) eating() {
	fmt.Printf("%s cat is eating\n", c.color)
}

func (d Dog) sleeping() {
	fmt.Printf("%s dog is sleeping\n", d.color)
}
func (d Dog) eating() {
	fmt.Printf("%s cat is eating\n", d.color)
}

func dosth(a Ainimal) {
	a.sleeping()
	a.eating()
}

func main() {
	c1 := Cat{"white"}
	d1 := Dog{"black"}
	c1.sleeping()
	d1.eating()
	//多态体验
	var a1 Ainimal
	a1 = c1
	a1.eating()
	a1 = d1
	a1.eating()
	fmt.Println("-----------------")
	dosth(a1)
	dosth(a1)
}
