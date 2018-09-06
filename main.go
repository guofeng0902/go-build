package main

import (
	"fmt"
	"net/http"
)

func main() {
    fmt.Println("=============hi======guofeng========hi==============")
	http.HandleFunc("/user/webhooks", func(writer http.ResponseWriter, request *http.Request) {
		request.ParseForm()
		writer.WriteHeader(200)
		fmt.Fprintf(writer, "收到webhook请求, 请求方法: %s, 请求参数: %v", request.Method, request.Form)
		fmt.Printf("收到webhook请求, 请求方法: %s, 请求参数: %v\n", request.Method, request.Form)
	})
	http.ListenAndServe("10.177.222.190:9000", nil)
}
