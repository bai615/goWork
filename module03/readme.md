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




## 构建多种系统架构支持的 Docker 镜像 -- docker manifest 命令详解

```shell
docker build -t bai615/arm64-go-httpserver:1.0.1 . 
docker build -t bai615/arm64v8-go-httpserver:1.0.1 . 
docker build -t bai615/amd64-go-httpserver:1.0.1 . 

```
创建 manifest 列表

# $ docker manifest create MANIFEST_LIST MANIFEST [MANIFEST...]
$ docker manifest create bai615/go-httpserver:1.0.1 \
bai615/amd64-go-httpserver:1.0.1 \
bai615/arm64-go-httpserver:1.0.1 \
bai615/arm64v8-go-httpserver:1.0.1

设置 manifest 列表

# $ docker manifest annotate [OPTIONS] MANIFEST_LIST MANIFEST
$ docker manifest annotate bai615/go-httpserver:1.0.1 \
bai615/arm64-go-httpserver:1.0.1 \
--os linux --arch arm64

$ docker manifest annotate bai615/go-httpserver:1.0.1 \
bai615/arm64v8-go-httpserver:1.0.1 \
--os linux --arch arm64 --variant v8

$ docker manifest annotate bai615/go-httpserver:1.0.1 \
bai615/amd64-go-httpserver:1.0.1 \
--os linux --arch amd64

### 查看 manifest 列表

$ docker manifest inspect bai615/go-httpserver:1.0.1

### 推送 manifest 列表
最后我们可以将其推送到 Docker Hub。


$ docker manifest push bai615/go-httpserver:1.0.1

### 参考：
[使用nsenter工具管理docker容器](https://www.jianshu.com/p/73a27f05dd82)
[构建多种系统架构支持的 Docker 镜像 -- docker manifest 命令详解](https://vuepress.mirror.docker-practice.com/image/manifest/)