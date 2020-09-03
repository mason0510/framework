package main

import (
  "fmt"
)

func main() {
  test_recover()
  fmt.Println("after recover")
}
//defer 会先处理写成再退出
//服务很容易宕机
func test_recover() {
  //defer func() {
  //  fmt.Println("defer func")
  //  //程序不会卦
  //  if err := recover(); err != nil {
  //    fmt.Println("recover success")
  //  }
  //}()
  //
  //arr := []int{1, 2, 3}
  //fmt.Println(arr[4])
  //fmt.Println("after panic")
  names := []string{"geektutu"}
  fmt.Println(names[0])
}
