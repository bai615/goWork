## 模块十的作业


### 1、为 HTTPServer 添加 0-2 秒的随机延时；
```go
func randInt(min int, max int) int {
    rand.Seed(time.Now().UTC().UnixNano())
    return min + rand.Intn(max-min)
}

func index(w http.ResponseWriter, r *http.Request) {
    delayMillisecs := randInt(0, 2000)
    delay := time.Millisecond * time.Duration(delayMillisecs)
    time.Sleep(delay)
	
    .......
		
}	  
```

### 2、为 HTTPServer 项目添加延时 Metric；
参见 metrics/metrics.go

### 3、将 HTTPServer 部署至测试集群，并完成 Prometheus 配置；

### 4、从 Promethus 界面中查询延时指标数据；

### 5、（可选）创建一个 Grafana Dashboard 展现延时分配情况。

提交地址： https://jinshuju.net/f/gFV4q6

截止日期：2023 年 2 月 12 日 23:59