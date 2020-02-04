# Step1: environment

## Download the Go distribution & Install Go tools

Download distribution from https://golang.org/dl/ or https://studygolang.com/dl

Install step follow https://golang.org/doc/install

*nix 平台存档包安装如下：
```
$ sudo  tar -C /usr/local -xzf go$VERSION.$OS-$ARCH.tar.gz
```
图形化安装包则直接点击安装.

Set new PATH environment variable for bin.
~/.bash_profile || /etc/bash.bashrc || /etc/profile || ~/.zshrc:
```
export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
```

Set workspace directory by setting GOPATH environment variable.
https://github.com/golang/go/wiki/SettingGOPATH
```
export GOPATH=$HOME/go
``````

Enforce it
```
$ source ~/.zshrc
```

Other go environment variable can also be set, see `go env`.

## A Global Proxy for Go Modules 

```
# Enable the go modules feature
export GO111MODULE=on

# Set the GOPROXY environment variable
export GOPROXY=https://goproxy.io
```

## Local Document

If godoc command is not in /usr/local/go/bin, run:
```
$ go get golang.org/x/tools/cmd/godoc
```

Then `godoc` will be install in $GOPATH/bin/godoc
```
$ godoc # http://localhost:6060/
```

## Test installation
```
package main

import "fmt"

func main() {
    fmt.Printf("hello, world\n")
}
```

```
$ go run hello.go  # 会根据 GOPATH 路径来找源文件
```

## Uninstall
```
https://golang.org/doc/install#uninstall
http://localhost:6060/doc/install#uninstall
```

---

# Step2: how to write go code
```
https://golang.org/doc/code.html
http://localhost:6060/doc/code.html
```

代码结构内容有 通常惯例、workspace、GOPATH、Import路径、第一个程序、第一个库、包名、测试、远程包。

## GOPATH 环境变量
```
GOPATH 环境变量用途是寻找 GO 代码，通过 `go env GOPATH` 查看，默认值为 $HOME/go

但是系统的 $GOPATH 是空的，如果要使用 $GOPATH，需要先执行 `export GOPATH=$(go env GOPATH)`   

自定义 GOPATH 的方法：https://golang.org/wiki/SettingGOPATH，
其实就是指定目录 `export GOPATH=/xxx/xxx`，( 如果当前所在目录就是go项目目录，那么就是 `export GOPATH=$(pwd)` )

workspace 的子目录 bin/ 出于便利可设置到 PATH 环境变量 `export PATH=$PATH:$(go env GOPATH)/bin`
```

## Import 路径
```
# import 路径作为一个包的唯一标识，建议使用个人账户相关作为基础路径，如 $GOPATH/src/github.com/username

$ export GOPATH=/Users/xxxx/go-lab
$ mkdir -p $GOPATH/src/github.com/farwish
$ mkdir $GOPATH/src/github.com/farwish/hello # 包目录
```

## 第一个程序
```
$ vi $GOPATH/src/github.com/farwish/hello/hello.go

# install 不加路径默认指向当前目录，会从 GOPATH 的 src 目录下找，install 会在 GOPATH 的 bin 目录生成可执行二进制文件 $GOPATH/bin/hello
$ go install github.com/farwish/hello
```

## 第一个库
```
# 包目录
$ mkdir $GOPATH/src/github.com/farwish/stringutil 

# build 不会生成文件，编译的包 保存在本地缓存中，可以在 hello.go 中 import 使用，然后重新 go install github.com/farwish/hello
# 要编译这个包用 go build github.com/farwish/stringutil
# 安装 hello 程序
$ go install github.com/farwish/hello
```

```
# 目录结构
bin/
    hello                 # command executable
src/
    github.com/farwish/
        hello/
            hello.go      # command source
        stringutil/
            reverse.go    # package source
```

## 包名
```
go 源文件的第一句必须是 `package xxxname`。

go 的惯例是 import 路径的最后一个元素作为包名, 所以包 "crypto/rot13" 的名字应为 rot13。

可执行命令必须总是使用包 main。

二进制中引用的所有包的包名不需要唯一，import 路径（包全名）唯一即可。
```

## 测试
```
Go 有一个由 go test 命令和 testing 包组成的轻量测试框架。

写测试通过创建一个名字 _test.go 结尾的文件，它包含签名为 func (t *testing.T) 的名字叫 TestXXX 的函数。
测试框架运行每一个这样的函数；如果函数调用了失败函数 t.Error 或 t.Fail，测试认为失败了。

添加一个测试到 stringutil 包，通过创建一个文件 $GOPATH/src/github.com/user/stringutil/reverse_test.go 包含如下代码：
```

```
package stringutil

import "testing"

func TestReverse(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	}
	for _, c := range cases {
		got := Reverse(c.in)
		if got != c.want {
			t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
```

运行测试：go test github.com/farwish/stringutil

一如既往，如果你在包所在目录使用 go 工具，可以忽略包路径，go test

go help test 查看更多细节。

## 远程包
```
# 拉取远程包，放到 GOPATH 指定的工作空间位置；如果包不存在，go get 跳过远程拉取并且和 go install 行为一样。
$ go get github.com/golang/example/hello
$ $GOPATH/bin/hello
```

在发出 go get 命令后，工作空间目录树如下：
```
bin/
    hello                           # command executable
src/
    github.com/golang/example/
	.git/                       	# Git repository metadata
        hello/
            hello.go                # command source
        stringutil/
            reverse.go              # package source
            reverse_test.go         # test source
    github.com/farwish/
        hello/
            hello.go                # command source
        stringutil/
            reverse.go              # package source
            reverse_test.go         # test source
```

# FAQ

## 如何下载 golang.org 的包

方式1.
```
# Go包管理器 https://github.com/gpmgo/gopm
$ go get -u -v github.com/gpmgo/gopm
$ ./bin/gopm get -u -v github.com/gin-gonic/gin # 不加-g参数会下载到$GOPATH位置的.vendor/下面
```

方式2.
每次都要指定完整路径, 比如官网包在github的镜像, go get -u github.com/golang/net/http（有依赖包时无法解决问题）


# 总结
```
export GOPATH=/home/wc/go-lab
go get -u -v github.com/gpmgo/gopm
export PATH=$PATH:/home/wc/go-lab/bin
gopm get -g -u -v github.com/gin-gonic/gin # -g 会安装到 GOPATH 目录，不加会安装到 .vendor 目录
```
