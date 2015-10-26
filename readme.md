# 0.Info
ryanemax/go-listen，使用原生net/http与io/ioutil库实现，用于监听HTTP端口的请求，同时将请求的详情记录在./log目录中。

- 模拟HTTP服务端，获取通过HTTP接入设备的所有请求

- 通过log文件可分析出该设备所有请求的格式及标准，便于开发出模拟服务端

# 1.Installation
go get github.com/ryanemax/go-listen
