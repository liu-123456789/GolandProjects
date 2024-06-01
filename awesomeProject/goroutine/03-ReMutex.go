package main

import (
	"fmt"
	"sync"
	"time"
)

var wg2 sync.WaitGroup
var rwlock sync.RWMutex
var begin = 100

/*
期望：读共享，写独占
使用场景，读多写少
func (r *rlocker) Lock()
func (r *rlocker) Unlock()
func (rw *RWMutex) RLock() ：申请读锁
func (rw *RWMutex) RUnlock() ：释放读锁
*/

func writer(num int) {
	for {
		rwlock.Lock()
		begin++
		fmt.Println("i am writer,begin,%d\n", num, begin)
		rwlock.Unlock() //释放锁
		time.Sleep(time.Millisecond * 3)
	}
	wg2.Done()
}
func reader(num int) {
	for {
		rwlock.RLock() //读取共享锁
		time.Sleep(time.Millisecond * 1)
		fmt.Println("i am reader,begin,%d\n", num, begin)
		rwlock.RUnlock() //释放锁
		time.Sleep(time.Millisecond * 3)
	}
	wg2.Done()
}

func main() {
	wg2.Add(10)
	for i := 0; i < 5; i++ {
		go reader(i)
	}
	for i := 0; i < 3; i++ {
		go writer(i)
	}
	wg2.Wait()
}
