package main

import "fmt"

func main() {
	var a1 [5]int = [5]int{1, 2, 3, 4, 5}

	s1 := a1[2:4]
	fmt.Println(s1)
	s1 = append(s1, 100, 1000)
	fmt.Println(s1)
	fmt.Println(a1)
}
