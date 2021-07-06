package main

import (
    "fmt"
)

func main() {
    // 初始化集合的方式
    countryCityMap := make(map[string]string)
    fmt.Printf("countryCityMap type is %T\n", countryCityMap)       // map[string]string
    fmt.Printf("countryCityMap value is %v\n", countryCityMap)      // map[]
    if (countryCityMap != nil) {
        fmt.Println("countryCityMap is not nil")
    }
    countryCityEnMap := map[string]string {"Italy": "Rome", "France": "Paris"}
    fmt.Printf("countryCityEnMap value is %v\n", countryCityEnMap)  // map[France:Paris Italy:Rome]
    fmt.Printf("countryCityEnMap length is %v\n", len(countryCityEnMap)) // 2

    // 插入键值
    countryCityMap["Italy"] = "罗马"
    countryCityMap["France"] = "巴黎"

    // 输出集合的键
    for countryEn := range countryCityMap {
        fmt.Println(countryEn)
    }

    // 查看集合中元素是否存在
    city, ok := countryCityMap["France"]
    fmt.Println(city) // 巴黎
    fmt.Println(ok)   // true

    icity, isok := countryCityMap["USA"]
    fmt.Println(icity) // ""
    fmt.Println(isok)  // false

    // 删除集合的元素 delete()
    delete(countryCityEnMap, "Italy")
    for countryEn := range countryCityEnMap {
        fmt.Println(countryEn) // France
    }
}
