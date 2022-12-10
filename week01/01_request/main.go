package main

import (
	"io"
	"strings"
)
import "net/http"

func param(res http.ResponseWriter, req *http.Request) {
	// text/plain; charset=utf-8
	res.Header().Add("Content-Type", "text/html; charset=utf-8")
	res.Header().Add("Access-Control-Allow-Origin", "*")
	res.Header().Add(
		"Access-Control-Allow-Methods",
		"OPTIONS, HEAD, GET, POST, PUT, DELETE",
	)
	//fmt.Fprintln(res, "第一个：接收客户端 request，并将 request 中带的 header 写入 response header")
	//req.Header中Header本质是:type Header map[string][]string
	header := req.Header
	//fmt.Fprintln(res, "Header 全部数据:", header)

	//明确给定类型
	//var acc []string =header["Accept"]
	//for _,n:=range acc{
	//	fmt.Fprintln(res,"Accepth内容:",n)
	//}

	// 自定义的头数据 abc
	//fmt.Fprintln(res, "Header abc数据:", header["Abc"])

	//fmt.Fprintln(res, "Header abc数据类型:", reflect.TypeOf(header["Abc"]))
	res.Header().Add("abc1", strings.Join(header["Abc"], ""))
	res.Header().Set("abc2", strings.Join(header["Abc"], ""))
	res.Header().Set("name1", "my name is smallsoup")
	res.Header().Add("name2", "my name is smallsoup")
	res.WriteHeader(404)
	res.Write([]byte("Hello World!"))
	io.WriteString(res, "200")
}

// 1、接收客户端 request，并将 request 中带的 header 写入 response header
func main() {
	server := http.Server{
		Addr: ":80",
	}
	http.HandleFunc("/param", param)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("name", "my name is smallsoup")
		w.WriteHeader(500)
		w.Write([]byte("hello world\n"))
	})
	server.ListenAndServe()
}
