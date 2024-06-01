package main

import (
	"fmt"
)

func main() {
	//int1 := []int{1, 6, 3}
	////fmt.Printf("%d", len(int1))
	//for i := len(int1) - 1; i >= 0; i-- {
	//	fmt.Println(int1[i], i)
	//}
	//fmt.Printf("%v", Itlists([]int{1, 2, 3, 4}, 9, 2))
	//Silces(4)

	//v, d := Silcet([]int{1, 2, 3, 4, 5}, 2)
	//fmt.Printf("%v,\n %d", v, d)
	v1, d1 := Silced([]int{1, 2, 3, 4, 5}, 2)
	fmt.Printf("删除后的 %v,\n删除的下标 %d", v1, d1)
}

//	func main2() {
//		dex := 4
//		var int1 = []int{1, 2, 3}
//		int1 = append(int1, dex)
//		for i := len(int1) - 1; i >= dex; i-- {
//			fmt.Println(int1[i])
//		}
//
// }
type Inlist interface {
	~int | int32 | int64
}

func Itlists[T any](vals []T, val T, dex int) []T {
	if dex < 0 || dex >= len(vals) {
		println("dex不可以用")
	}
	vals = append(vals, val)
	for i := len(vals) - 1; i > dex; i-- {
		if i-1 >= 0 {
			vals[i] = vals[i-1]

		}
	}
	vals[dex] = val
	return vals
}

func Silces(dex int) {
	silces1 := []int{6, 5, 4, 9, 8, 7}

	silces1 = silces1[:dex]
	for k, v := range silces1 {
		println(k, v)
	}
	fmt.Printf("\n %d, %d", len(silces1), cap(silces1))
}

func Silcet[T any](val []T, dex int) ([]T, int) {
	if dex >= len(val) || dex < 0 {
		println("输入不合法")
	}
	val = val[:dex]

	return val, dex
}

type Silcein interface {
	int
}

func Silced[T Silcein](val []T, dex int) ([]T, int) {
	vala := []T{}
	if dex >= len(val) || dex < 0 {
		println("输入不合法")
	}
	for k, v := range val {
		if k != dex-1 {
			vala = append(vala, v)

		}

	}
	return vala, dex

}
