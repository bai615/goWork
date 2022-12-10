package main

import (
	"io"
	"log"
	"net/http"
)

// 定义 handle 处理函数
func healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Hello World!"))
	io.WriteString(w, "200")
}

// 4、当访问 localhost/healthz 时，应返回 200
func main() {

	// 注册 handle 处理函数
	http.HandleFunc("/healthz", healthz) //Use the default DefaultServeMux.
	// ListenAndService
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}
