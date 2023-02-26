## 模块十二作业


把我们的 httpserver 服务以 Istio Ingress Gateway 的形式发布出来。以下是你需要考虑的几点：

### 1、如何实现安全保证；

### 2、七层路由规则；

### 3、考虑 open tracing 的接入。


提交地址：  https://jinshuju.net/f/LwsTE6

截止日期：2023 年 2 月 26 日 23:59


### 安装 istio
```shell
curl -L https://istio.io/downloadIstio | sh -

# 检查一下 istio 环境以及版本信息：
$ istioctl version
client version: 1.12.1
control plane version: 1.12.1
data plane version: 1.12.1 (2 proxies)

# 对其进行验证：
$ istioctl verify-install -f my-demo-config.yaml
......
Checked 14 custom resource definitions
Checked 3 Istio Deployments
✔ Istio is installed and verified successfully

$ kubectl create ns istio-app
namespace/istio-app created
$ kubectl label ns istio-app istio-injection=enabled
namespace/istio-app labeled
```


## 链路追踪 Tracing

安装 Jaeger
```
kubectl apply -f tracing-jaeger.yaml
```

配置采样比例
```
kubectl edit configmap istio -n istio-system

```

```
apiVersion: v1
data:
  mesh: |-
    accessLogFile: /dev/stdout
    defaultConfig:
      discoveryAddress: istiod.istio-system.svc:15012
      proxyMetadata: {}
      tracing:
        sampling: 100
        zipkin:
          address: zipkin.istio-system:9411
    enablePrometheusMerge: true
```
