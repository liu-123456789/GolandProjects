package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	go func() {
		for {
			num, ok := <-ch1
			fmt.Println("ok=", ok)
			if !ok {
				break
			}
			ch2 <- num * num
		}
		close(ch2)
	}()

	for {
		msg, ok := <-ch2
		if !ok {
			break
		}
		fmt.Println(msg)
	}

}
