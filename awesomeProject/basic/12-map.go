package main

import (
	"fmt"
)

func main() {
	var scores map[string]uint

	//向map添加元素
	//map为空不能添加元素
	//map再使用前必须make
	scores = make(map[string]uint)
	scores["q"] = 99
	scores["w"] = 98
	fmt.Println(scores)
	fmt.Println("q is score is", scores["q"])
	val, ok := scores["e"]
	fmt.Println(val, ok)

	//map 遍历
	fmt.Println(len(scores))
	//使用一个range可以遍历map，遍历所有的容器

	for k, v := range scores {
		fmt.Println("scores[%s] = %d\n", k, v)
	}
	a := [10]int{1, 2, 3, 4, 5, 6}
	for a1, a2 := range a {
		fmt.Println(a1, a2)
	}
}
