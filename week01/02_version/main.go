package main

import (
	"net/http"
	"os"
)

// 2、读取当前系统的环境变量中的 VERSION 配置，并写入 response header
func main() {
	// VERSION=1.0.1 go run main.go
	//version := os.Getenv("VERSION")
	//fmt.Println("version is:", version)

	server := http.Server{
		Addr: ":80",
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// GO_VERSION_DEMO=1.0.2
		version := os.Getenv("GO_VERSION_DEMO")
		w.Header().Set("version", version)
		w.Write([]byte("get version demo.\n"))
	})
	server.ListenAndServe()
}
