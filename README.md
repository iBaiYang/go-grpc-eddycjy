# go-grpc-eddycjy

## 引言

无意中发现了这本书，[Go 语言编程之旅](https://golang2.eddycjy.com/)，
这本书中的 [示例代码](https://github.com/go-programming-tour-book) 地址也找到了。

书的作者 煎鱼 的博客地址是 [煎鱼博客](https://eddycjy.com/) ，
博客的 [Github地址](https://github.com/eddycjy/blog) ，
作者的 [Github地址](https://github.com/eddycjy)。

里面看到了 [Gin框架示例](https://github.com/eddycjy/go-gin-example)，
相应的一系列 [博文](https://github.com/EDDYCJY/go-gin-example/blob/master/README_ZH.md)，
里面说是 [Gin实践](https://segmentfault.com/a/1190000013297625) 的连载，在 segmentfault 网站上，
另外在作者博客中也记录了这个系列，不过内容看着不一样， [gin](https://eddycjy.com/tags/gin/)，
看着内容是作者博客 [Go语言入门系列](https://eddycjy.com/go-categories/) 中的一部分。
对比看了下，segmentfault 上那个是2018年写的，好多已经过时了，就不要看了。

作者还有 [Go 语言进阶之旅](https://golang1.eddycjy.com/)。

这是跟着作者学习的一个示例项目。

## 准备

win10下，在github <https://github.com/protocolbuffers/protobuf/releases/ >下载 Protocol Buffers 。
解压后，设置环境变量：`G:\WorkSoft\protoc\bin` 。重启，查看版本(libprotoc 23.4)：
```
> protoc --version
```

protoc 插件安装：
> go get -u github.com/golang/protobuf/protoc-gen-go@v1.3.2

到目录 `G:\GoPath\pkg\mod\github.com\golang\protobuf@v1.3.2\protoc-gen-go` 下执行 `go build .`，
把生成的 `protoc-gen-go.exe` 移动到目录 `G:\GoPath\bin` 下。

## 命令

### 获取包

protoc 插件安装：
> go get -u github.com/golang/protobuf/protoc-gen-go@v1.3.2

gRPC 库安装：
> go get -u google.golang.org/grpc@v1.29.1

### 实操

生成 proto 文件：
```
> protoc --go_out=plugins=grpc:. ./proto/helloworld.proto
```
#### Unary RPC：一元 RPC

```
> protoc --go_out=plugins=grpc:. ./proto/unaryrpc.proto
```

> go run server.go

> go run client.go

#### Server-side streaming RPC：服务端流式 RPC

```
> protoc --go_out=plugins=grpc:. ./proto/serversidestreamingrpc.proto
```





