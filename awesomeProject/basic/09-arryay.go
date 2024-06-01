package main

import "fmt"

func main() {
	fmt.Println(val)
	fmt.Println("the 3rd element is", val[3])
	fmt.Println(a2)
	fmt.Println("------------")
	fmt.Println(a2[2][1])
}

var val [5]int = [5]int{1, 2, 3, 4, 5}

var a2 [3][4]int = [3][4]int{
	{1, 2, 3, 4},
	{11, 22, 33, 44},
	{21, 22, 23, 24},
}
