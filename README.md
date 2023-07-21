# go-grpc-eddycjy

## 引言

无意中发现了这本书，[Go 语言编程之旅](https://golang2.eddycjy.com/)，
这本书中的 [示例代码](https://github.com/go-programming-tour-book) 地址也找到了。

书的作者 煎鱼 的博客地址是 [煎鱼博客](https://eddycjy.com/) ，
博客的 [Github地址](https://github.com/eddycjy/blog) ，
作者的 [Github地址](https://github.com/eddycjy)。

里面看到了 [gRPC 和 Protobuf] 的相关内容。

RPC 代指远程过程调用（Remote Procedure Call），它的调用包含了传输协议和编码（对象序列）协议等等，
允许运行于一台计算机的程序调用另一台计算机的子程序，而开发人员无需额外地为这个交互作用编程，
因此我们也常常称 RPC 调用，就像在进行本地函数调用一样方便。

gRPC 是一个高性能、开源和通用的 RPC 框架，面向移动和基于 HTTP/2 设计。目前提供 C、Java 和 Go 语言等等版本，
分别是：grpc、grpc-java、grpc-go，其中 C 版本支持 C、C++、Node.js、Python、Ruby、Objective-C、PHP 和 C# 支持。
gRPC 基于 HTTP/2 标准设计，带来诸如双向流、流控、头部压缩、单 TCP 连接上的多复用请求等特性。
这些特性使得其在移动设备上表现更好，在一定的情况下更节省空间占用。
gRPC 的接口描述语言（Interface description language，缩写 IDL）使用的是 Protobuf，都是由 Google 开源的。

Protocol Buffers（Protobuf）是一种与语言、平台无关，可扩展的序列化结构化数据的数据描述语言，
我们常常称其为 IDL，常用于通信协议，数据存储等等，相较于 JSON、XML，它更小、更快。

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

#### Client-side streaming RPC：客户端流式 RPC

```
> protoc --go_out=plugins=grpc:. ./proto/clientsidestreamingrpc.proto
```

#### Bidirectional streaming RPC：双向流式 RPC

```
> protoc --go_out=plugins=grpc:. ./proto/bidirectionalstreamingrpc.proto
```
