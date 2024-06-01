package main

import "fmt"

func main() {
	var a1 []int
	for i := 0; i <= 50; i++ {
		a1 = append(a1, i)
		fmt.Println("len=", len(a1), "cap=", cap(a1))
	}
}
