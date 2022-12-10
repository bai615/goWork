package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// 3、Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
func main() {

	server := http.Server{
		Addr: ":80",
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		var logMap map[string]string /*创建集合 */
		logMap = make(map[string]string)

		//fmt.Fprintln(w, "Header 全部数据:", w.Header())

		logMap["time"] = time.Now().Format("2006-01-02 15:04:05")
		logMap["ip"] = r.RemoteAddr
		logMap["method"] = r.Method
		logMap["statusCode"] = w.Header().Get("statusCode")
		logMap["path"] = r.RequestURI
		logMap["query"] = r.URL.RawQuery

		logStr, _ := json.Marshal(logMap)

		fmt.Printf("%s\n", logStr)

		w.Write([]byte("hello world\n"))
	})
	server.ListenAndServe()

	/*
		{"ip":"[::1]:62418","method":"GET","path":"/healthz","query":"","statusCode":"","time":"2022-12-10 19:43:08"}
		{"ip":"[::1]:62418","method":"GET","path":"/healthz","query":"","statusCode":"","time":"2022-12-10 19:43:09"}
		{"ip":"[::1]:62418","method":"GET","path":"/healthz?id=132","query":"id=132","statusCode":"","time":"2022-12-10 19:43:38"}
		{"ip":"[::1]:62418","method":"GET","path":"/healthz?id=132","query":"id=132","statusCode":"","time":"2022-12-10 19:43:39"}

	*/
}
