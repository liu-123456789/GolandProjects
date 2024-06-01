package main

import (
	"fmt"
	"time"
)

// 传递字符串类型的通道
var ch chan string

func dosthAfterReadChannel() {
	fmt.Println("开始读取")
	//time.Sleep(time.Second * 5)
	msg := <-ch
	fmt.Printf("读取成功:", msg)
}

func main() {
	//使用make 对ch进行初始化
	ch = make(chan string) //无缓冲区
	go dosthAfterReadChannel()
	fmt.Println("开始写入")
	time.Sleep(time.Second * 5)
	//写channel
	ch <- "hello"
	fmt.Println("写入完成")
}
