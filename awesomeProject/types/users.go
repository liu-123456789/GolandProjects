package main

import (
	"fmt"
)

func ChangeUser() {
	u1 := User{name: "tom", age: 18}
	fmt.Printf("%+v \n", u1)
	fmt.Printf("u1 address %p \n", u1)

	u1.changeName("jerry")
	u1.changeage(35)
	fmt.Printf("%+v", u1)
	println("=======")
	u2 := &User{name: "tom2", age: 19}
	fmt.Printf("%+v \n", u2)
	fmt.Printf("U2 address %p \n", u2)
	u2.changeage(22)
	u2.changeName("jerry2")
	fmt.Printf("%+v", u2)
}

type User struct {
	age      int
	name     string
	nikeName string
}

func (u User) changeName(name string) {
	u.nikeName = name
}

func (u *User) changeage(age int) {
	u.age = age
}

type LinkedList struct {
	head node
}

func (l LinkedList) Add(index int, val any) error {
	//TODO implement me
	panic("implement me")
}

func (l LinkedList) Update(index int, val any) error {
	//TODO implement me
	panic("implement me")
}

func (l LinkedList) Delect(index int, val any) error {
	//TODO implement me
	panic("implement me")
}

type node struct {
	next *node
}
