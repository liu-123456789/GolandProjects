package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	//1.绑定ip端口,设置侦听
	//func Listen(network,address string)(Listener,error)
	//net.Listen("tcp","127.0.0.1:8080") //只能本机访问
	listener, err := net.Listen("tcp", ":8080") //0.0.0.0 本机任意有效ip都可以昨晚访问ip
	if err != nil {
		log.Print("failed to Listener", err)
	}
	fmt.Println("local addr", listener.Addr())
	//4.扫尾工作
	defer listener.Close()

	//2.等待客户端链接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("failed to accept", err)
			continue
		}
		//3.与客户端通信 使用conn
		//启动一个goroutine去和客户端通信
		go handle_conn(conn)
	}

}

func handle_conn(conn net.Conn) {
	defer conn.Close()
	fmt.Println("remoter addr", conn.RemoteAddr().String())
	//循环通信
	buf := make([]byte, 256)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("failed to read", err)
			break
		}
		fmt.Printf("read %d bytes,mag(%s)\n", n, string(buf))
		//写回给网络
		conn.Write(buf[:n])
	}

}
