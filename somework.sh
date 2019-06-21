#!/bin/bash
#
# Set $GOPATH first at the beginning.
#
# @author farwish <farwish@foxmail.com>


/etc/bash.bashrc 加入:
```
export GOPATH=/home/wc/go-lab
export GO111MODULE=on
export GOPROXY=https://goproxy.io
export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin:/usr/local/protobuf/bin
```

注意：
开启GO mod之后，下载的包将存放在 $GOPATH/pgk/mod 路径, GOPATH (GOPATH/src下的包) 不再用于解析 imports 包路径.
$GOPATH/bin 路径的功能依旧保持

micro 工具
```
go get -u github.com/micro/micro
```

micro 依赖项 consul, protobuf

依赖(consul):
```
$ wget https://releases.hashicorp.com/consul/1.5.1/consul_1.5.1_linux_amd64.zip
$ unzip consul_1.5.1_linux_amd64.zip
$ sudo mv consul /usr/local/bin/
$ consul agent -dev  # 开发模式启动agent
$ # 默认注册中心是 mdns, 使生效：
$ export MICRO_REGISTRY=consul 或者 micro --registry=consul xxx
```

依赖(protoc, protoc-gen-go):
1> Protocol Compiler 安装：https://github.com/protocolbuffers/protobuf
    (https://github.com/protocolbuffers/protobuf/blob/master/src/README.md)
    ```
    $ sudo apt-get install autoconf automake libtool curl make g++ unzip
    $ wget https://github.com/protocolbuffers/protobuf/releases/download/v3.8.0/protobuf-all-3.8.0.tar.gz
    $ tar zxf protobuf-all-3.8.0.tar.gz && cd protobuf-3.8.0
    $ ./configure --prefix=/usr/local/protobuf # 自定义安装目录，记得把 /usr/local/protobuf/bin 加入到PAT
H环境变量以使用 protoc 命令
    $ make
    $ make check # 检查所有的特性在当前系统上是否都支持
    $ sudo make install
    $ sudo ldconfig # 刷新共享库缓存
    ```

2> Protbuf Runtime 安装：https://github.com/golang/protobuf
    (https://github.com/golang/protobuf#installation)
    ```
    $ go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
    ```
    使用方法：https://github.com/golang/protobuf#using-protocol-buffers-with-go

3> micro 的 protobuf 插件(代码生成器)安装：
    ```
    # 不加 -u, 处理 go get: error loading module requirements
    # https://stackoverflow.com/questions/55430150/go111module-on-error-loading-module-requirements
    $ go get github.com/micro/protoc-gen-micro
    ```

编写服务
    ```
    $ micro new github.com/micro/example
    $ cd /home/wc/go-lab/src/github.com/micro/example
    $ protoc --proto_path=.:$GOPATH/src --go_out=. --micro_out=. proto/example/example.proto
    $ go run main.go
    ```


FAQ

1.go get: error loading module requirements
go get 不加 -u，不升级依赖

2.$GOPATH/go.mod exists but should not
开启 go 模块后(GO111MODULE=on), go.mod 与 项目不能在 GOPATH 中共存.

