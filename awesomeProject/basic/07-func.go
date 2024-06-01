package main

import "fmt"

func main() {
	x, y := 10, 20
	fmt.Println(add(x, y))
	f := add(x, y)
	fmt.Println(f)
	fmt.Println("-----------------")
	f1 := addorsub(2, 5, add)
	f2 := addorsub(10, 5, sub)
	fmt.Println(f1, f2)
	println("-----------------")
	val := func(a, b int) int {
		return a * b
	}(10, 20)
	fmt.Println(val)
}

func add(a int, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func addorsub(a, b int, f func(x, y int) int) int {
	return f(a, b)
}
