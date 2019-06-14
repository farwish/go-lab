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

micro
```
go get -u github.com/micro/micro
```

micro 依赖 consul, protobuf

依赖(consul):
```
$ wget https://releases.hashicorp.com/consul/1.5.1/consul_1.5.1_linux_amd64.zip
$ unzip consul_1.5.1_linux_amd64.zip
$ sudo mv consul /usr/local/bin/
$ consul agent -dev  # 启动开发
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

3> go-micro protoc 生成器安装：
    ```
    $ go get -u github.com/micro/protoc-gen-micro
    ```


