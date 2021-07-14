# go-web-example
 一个基于 Go 语言的 Web 项目示例

## 构建项目
```shell
# sonic: 外部引用的json序列化方式
# nomsgpack: 禁用Gin的默认渲染
# mod=vendor：模块通过项目下的 vendor 文件夹构建
go build -tags=sonic,nomsgpack .
```

## 性能

1. `./main`
```text
Running 1m test @ http://localhost:8080/ping
  16 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     4.60ms    5.47ms 190.17ms   95.45%
    Req/Sec     1.44k   192.11     1.94k    74.78%
  1373029 requests in 1.00m, 191.18MB read
Requests/sec:  22863.79
Transfer/sec:      3.18MB
```

2. `./main -appMode prod`
```text
Running 1m test @ http://localhost:8080/ping
  16 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.56ms    1.07ms  32.77ms   63.97%
    Req/Sec     3.96k   463.79    33.82k    85.92%
  3779812 requests in 1.00m, 526.29MB read
Requests/sec:  62892.69
Transfer/sec:      8.76MB
```
