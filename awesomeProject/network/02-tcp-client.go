package main

import (
	"log"
	"net"
	"os"
)

func main() {
	//1.链接服务器
	conn, errs := net.Dial("tcp", "127.0.0.1:8080")
	if errs != nil {
		log.Print("failed to Dial", errs)
	}
	defer conn.Close()
	//2.通信 模仿nc命令：接受标准输入，发送给服务器
	for {
		buf1 := make([]byte, 256)
		//读取标准输入 --发送网络
		n, err := os.Stdin.Read(buf1)
		if err != nil {
			log.Print("failed to Dial", err)
		}
		if n > 0 {
			conn.Write(buf1)
		}
		//读完网络 --打印
		n1, err1 := conn.Read(buf1)
		if err != nil {
			log.Print("failed to Dial", err1)
		}
		os.Stdin.Write(buf1[:n1])
	}
}
