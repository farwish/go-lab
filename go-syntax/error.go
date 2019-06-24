package main

import (
    "fmt"
    "errors"
)
/* error接口类型
type error interface {
    Error() string
}
*/

func main() {
    result, err := Sqrt(-1)

    if err != nil {
        fmt.Println(err.Error())
    } else {
        fmt.Println(result)
    }
}

func Sqrt(f float64) (float64, error) {
    if f < 0 {
        return 0, errors.New("error:f < 0")
    }

    return 0, nil
}
