# api

## Golang 安装、环境变量、模块初始化
    https://github.com/farwish/go-lab
    https://github.com/farwish/go-lab/blob/master/somework
    新项目go.mod初始化示例: go mod init backer

## Gin Web Framework
    https://github.com/gin-gonic/gin
    https://gin-gonic.com/zh-cn/docs/quickstart/

    路由组
    https://gin-gonic.com/zh-cn/docs/examples/grouping-routes/

    路由参数
    https://gin-gonic.com/zh-cn/docs/examples/param-in-path/

    query和post参数
    https://gin-gonic.com/zh-cn/docs/examples/query-and-post-form/

## ini 配置解析
    https://github.com/go-ini/ini
    https://ini.unknown.io

## GORM
    https://github.com/jinzhu/gorm
    https://gorm.io/zh_CN/docs/
    http://gorm.book.jasperxu.com/models.html#md
    ```
    $ go get -u github.com/jinzhu/gorm
    ```

## Swag 接口文档
    ```
    $ go get -u github.com/swaggo/swag/cmd/swag@v1.6.0
    $ swag -v
    $ go get -u github.com/swaggo/gin-swagger@v1.1.0
    $ go get -u github.com/swaggo/gin-swagger/swaggerFiles@v1.1.0
    ```

## 热重载
    ```
    $ go get github.com/pilu/fresh
    $ cd $GOPATH/xxxx
    $ fresh
    ```

## 采集库
    https://github.com/PuerkitoBio/goquery


[main 函数 与 init 函数]
main 函数只能在 package main 中，必须有；
init 函数可在 package main 和其它包中，是可选的；
这两个函数不能有参数和返回值。

[导入包的方式]
标准包使用给定的段路径，如 import "net/http"
第三方包通过指定路径，如 import "github.com/farwish/hello"
别名导入，使用时可以直接使用别名，如 import h "backer/handler"
点号导入，使用的时候就可以忽略报名，如 import . "backer/handler"
下划线导入，导入都会执行包的 init 函数，如果不需要使用这些包时采用，如 import _ "backer/handler"

[GORM]
模型定义的结构体标记:https://gorm.io/zh_CN/docs/models.html
模型定义的字段首字母大写，小写会被忽略。
MySQL的DATE/DATATIME类型可以对应Golang的time.Time
设置表属性：db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8 AUTO_INCREMENT=1;").AutoMigrate(&User{})
smallint型用法：https://www.codercto.com/a/43809.html
当前的datetime：time.Now().Format("2006-01-02 15:04:05")
随机数使用 math/rand 包：rand.Seed(time.Now().UnixNano()); rand.Intn(100)
int类型转换string使用 strconv 包：strconv.Itoa(1000)
查询指定列？

Gin+GORM+Restful
https://www.jianshu.com/p/35665b584347
https://www.jianshu.com/p/443766f0e796
https://www.jianshu.com/p/b34c3f17b417

代码查看平台:
    https://sourcegraph.com/github.com/gin-gonic/gin

database/sql: 标准库
