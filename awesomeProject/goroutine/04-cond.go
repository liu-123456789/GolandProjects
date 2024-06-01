package main

import (
	"fmt"
	"sync"
	"time"
)

var cond *sync.Cond
var lock01 sync.Mutex
var wg3 sync.WaitGroup

var beginnum = 1000
var prodictes []int
var custIdx = 0 //消费者读取下标位置

/*

func (c *Cond) Wait()
func (c *Cond) Signal()
func (c *Cond) Broadcast()
func NewCond(l Locker)
*/

/*
生产者模拟生产线，生产商品，整形切片，记录商品编号
消费者使用商品，读取出商品编号
*/
func productor() {
	for {
		lock01.Lock()
		prodictes = append(prodictes, beginnum)
		fmt.Printf("I am productor %d\n", beginnum)
		beginnum++
		time.Sleep(time.Millisecond * 2)
		lock01.Unlock()
		cond.Signal()
		time.Sleep(time.Millisecond * 3)
	}
}

func customer(num int) {
	for {

		//抢到锁
		lock01.Lock()

		//wait
		cond.Wait()
		if len(prodictes) == custIdx {
			//没有新的共享资源
			lock01.Unlock()
			continue
		}
		fmt.Printf("I am %d customer,product %d\n", num, prodictes[custIdx])
		custIdx++
		time.Sleep(time.Millisecond * 2)
		//释放锁
		lock01.Unlock()
		time.Sleep(time.Millisecond * 2)
	}
}

func main() {
	cond = sync.NewCond(&lock01)
	wg3.Add(3)
	go customer(1)
	go customer(2)
	time.Sleep(time.Second * 2)
	go productor()
	wg3.Wait()
}
