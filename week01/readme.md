
## 本周作业

编写一个 HTTP 服务器，大家视个人不同情况决定完成到哪个环节，但尽量把 1 都做完：

### 1、接收客户端 request，并将 request 中带的 header 写入 response header
[程序入口](01_request/main.go)

### 2、读取当前系统的环境变量中的 VERSION 配置，并写入 response header
[程序入口](02_version/main.go)

### 3、Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
[程序入口](01_request/main.go)

### 4、当访问 localhost/healthz 时，应返回 200
[程序入口](04_healthz/main.go)