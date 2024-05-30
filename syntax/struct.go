package main

import (
    "fmt"
)

// 定义类型名 Book 的结构, type + struct
type Book struct {
    title string
    author string
    year int
    id int
}

type Ebook struct {
    // 继承父结构元素
    Book
    url string
}

func main() {
    // 没有 key 的情况按顺序不变
    mayun := Book {"这就是马云", "陈", 2012, 1}
    fmt.Println(mayun) // {这就是马云 陈 2012 1}

    // 有 key 的情况顺序可变, 打印出来的顺序不会变
    xinbake := Book {author: "老外", title: "星巴克传奇", year: 2012, id: 2}
    fmt.Println(xinbake)

    // 忽略字段的值为 空字符 和 0
    jinyichuangye := Book {title: "精益创业"} // {精益创业  0 0}
    fmt.Println(jinyichuangye)

    // 结构字段赋值
    xinbake.author = "列夫"
    fmt.Println(xinbake)

    // 结构体作为函数参数
    printBook(mayun)

    // 结构体指针,
    var xiaomayun *Book
    xiaomayun = &mayun
    (*xiaomayun).title = "这就是小马云"
    // 修改一个之后两个都变
    printBook(*xiaomayun)
    printBook(mayun)

    // 通过结构的方法设置值
    mayun.setBookName("这就是小马云2")
    printBook(*xiaomayun)

    // 函数传指针的方式修改
    setBookName(&mayun, "这就是小马云3")
    printBook(mayun)

    // 通过结构的方法取值
    fmt.Println("通过结构的方法取值:" + mayun.getBookName())
}

func printBook(mybook Book) {
    fmt.Printf("title=%s\n", mybook.title)
    fmt.Printf("author=%s\n", mybook.author)
    fmt.Printf("year=%d\n", mybook.year)
    fmt.Printf("id=%d\n", mybook.id)
}

// 传指针的方式修改元素
func setBookName(b *Book, title string) {
    b.title = title
}

// 对象方法的方式
func (b *Book) setBookName(title string) {
    b.title = title
}

func (b *Book) getBookName() string {
    return b.title
}
