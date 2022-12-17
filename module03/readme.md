## 本周作业

### 构建本地镜像
参见Dockerfile文件，进行两阶段构建。

### 编写 Dockerfile 将模块二作业编写的 httpserver 容器化
```shell
 docker build -t bai615/go-httpserver:1.0.0 . 
 GOOS=linux CGO_ENABLED=0 GOARCH=amd64 docker build -t bai615/go-httpserver:1.0.0 . 
```

### 将镜像推送至 docker 官方镜像仓库
```shell
 docker push bai615/go-httpserver:1.0.0
```

### 通过 docker 命令本地启动 httpserver
```shell
 docker run --name mygo-httpserver -p 8080:8080 -d bai615/go-httpserver:1.0.0
```
### 通过 nsenter 进入容器查看 IP 配置
```shell
# 输入命令查看容器 <container id> 的进程的PID
docker inspect -f {{.State.Pid}}  <container id>
# 连接到这个容器，格式:  
nsenter --target    $PID   --mount --uts   --ipc   --net   --pid
```
### 作业需编写并提交 Dockerfile 及源代码。



### 参考：
[使用nsenter工具管理docker容器](https://www.jianshu.com/p/73a27f05dd82)