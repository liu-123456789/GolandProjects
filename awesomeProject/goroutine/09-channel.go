package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int, 6)

	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- i
			fmt.Println("次数+", i)
		}
		close(ch1)
	}()
	for {
		mag, ok := <-ch1
		if !ok {
			break
		}
		fmt.Println(mag)
		time.Sleep(time.Second * 1)
	}

}
