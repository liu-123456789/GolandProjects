package main

import (
	"fmt"
	"time"
)

func main() {
	/*
			定时器2类
		一次性
		周期性
	*/
	//一次性
	tm1 := time.NewTimer(time.Second * 3)
	msg1 := <-tm1.C
	fmt.Println(msg1)

	//周期性
	tk1 := time.NewTicker(time.Second * 2)
	for {
		msg2 := <-tk1.C
		fmt.Println(msg2)
	}
}
