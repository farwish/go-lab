package main

import (
    "fmt"
)

func main() {
    // 初始化集合的方式
    countryCityMap := make(map[string]string)
    countryCityEnMap := map[string]string {"Italy": "Rome", "France": "Paris"}

    // 插入键值
    countryCityMap["Italy"] = "罗马"
    countryCityMap["France"] = "巴黎"

    // 输出集合的值
    for country := range countryCityMap {
        fmt.Println(country)
    }

    // 查看集合中元素是否存在
    city, ok := countryCityMap["France"]
    fmt.Println(city) // 巴黎
    fmt.Println(ok)   // true

    icity, iok := countryCityMap["USA"]
    fmt.Println(icity) // ""
    fmt.Println(iok)   // false

    // 删除集合的元素 delete()
    delete(countryCityEnMap, "Italy")
    for countryEn := range countryCityEnMap {
        fmt.Println(countryEn) // France
    }
}
