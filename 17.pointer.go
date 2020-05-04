package main

import "fmt"

// 操作符&指的是变量地址（即获取指针地址）
// 操作符*指的是从指针地址中读取目标对象
func main () {
  value := 1
  fmt.Println("initial:", value)
  // 指针地址
  fmt.Println("pointer:", &value)
  // 变量实体
  pointer := &value
  fmt.Println("value:", *pointer)
}
