package main

import "fmt"

func main() {
	//DeferClosureLoopV1()
	//DeferClosureLoopv2()
	DeferClosureLoopV3()
	//forc()
}

func DeferClosureLoopV1() {
	for i := 0; i < 10; i++ {
		defer func() {
			println(i)

		}()
	}
}

func DeferClosureLoopv2() {
	for i := 0; i < 10; i++ {
		{
			defer func(i int) {
				println(i)
			}(i)
		}
	}

}
func forc() {
	for i := 0; i < 10; i++ {
		println(i)
		addr := i
		fmt.Printf("变量地址", addr)
	}
}

func DeferClosureLoopV3() {
	for i := 0; i < 10; i++ {
		j := i
		defer func() {
			println(j)
		}()
	}

}
