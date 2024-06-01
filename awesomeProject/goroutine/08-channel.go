package main

import "fmt"

func main() {
	che1 := make(chan int)
	che2 := make(chan int)
	go func(ch chan<- int) {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}(che1)

	go func(rch <-chan int, wch chan<- int) {
		for {
			num, ok := <-rch
			if !ok {
				break
			}
			wch <- num * num
		}
		close(wch)
	}(che1, che2)

	for {
		msg, ok := <-che2
		if !ok {
			break
		}
		fmt.Println(msg)
	}
}
