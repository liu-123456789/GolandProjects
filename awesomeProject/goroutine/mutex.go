package main

import (
	"fmt"
	"sync"
)

var val = 0
var wg1 sync.WaitGroup
var lock sync.Mutex //锁

func add() {
	lock.Lock() //锁住资源
	val++
	lock.Unlock() //释放资源
	wg1.Done()
}

func main() {
	wg1.Add(1000)
	for i := 0; i < 1000; i++ {
		go add()
	}
	wg1.Wait()
	fmt.Println("val 的次数", val)
}
