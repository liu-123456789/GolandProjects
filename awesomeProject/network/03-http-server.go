package main

import "net/http"

func HandleFunc(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("hello world"))
}

func ByeServer(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("bye bye"))
}

func main() {
	//设置路由规则
	//HandleFunc(pattern string, handler func(ResponseWriter, *Request))
	http.HandleFunc("/hello", HandleFunc)
	http.HandleFunc("/bye", ByeServer)

	//启动http服务
	http.ListenAndServe(":9090", nil)
}
