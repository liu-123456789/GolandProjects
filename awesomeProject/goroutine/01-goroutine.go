package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

/*
WaitGroup 的三个方法
func (wg *WaitGroup) Add(delta int) 设置要监控的goroutine
func (wg *WaitGroup) Done() 减少要监控的数值
func (wg *WaitGroup) Wait() 阻塞和等待监控的数值为0
*/

func dousth() {
	fmt.Println("I am a goroutine ,do sth")
	time.Sleep(time.Second * 5)
	fmt.Println("I am a goroutine ,do sth end")
	wg.Done() //减少一个监控记录
}

func main() {
	wg.Add(1)
	go dousth()
	wg.Wait()
	//dousth()
	fmt.Println("I am main goroutine")
	//time.Sleep(time.Second * 1)
}
