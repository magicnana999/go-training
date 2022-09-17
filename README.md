# 题1 我选WebSocket 端口8080
 
## 实现方法


* 基于gnet实现，WS的解析来自于gobwas
* Gin也可以实现，原理也差不多一样，

## 关于3点要求的说明

1.简单API设计说明文档， 罗列所实现的API路径及调用说明。 以README.md文档形式即可
>没太理解什么意思，这个WebSocket没有任何的业务方法，无论上行什么都会返回hello，没有业务的API。

## 运行说明

- 运行server_text.go 启动服务端
- 客户端我安装了mac版本的WebSocket Client 



# 题2
- 我本机是的Golang 是1.18版本，为了避免和您的机器上版本不一致，所以build的过程也放到镜像中构建，构建完成后直接启动。当然也可以本地构建。
- docker build 命令
```go
docker build --build-arg SV=1023 -t jinsong/go-training:1.0.7 .
//SERVER_VERSION 通过SV指定
```
- docker run 命令
```go
docker run -p 8080:8080 jinsong/go-training:1.0.7     
docker run -p 8080:8080 -d jinsong/go-training:1.0.7
```





