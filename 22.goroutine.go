package main

import "fmt"

func test1 (value string) {
  for i := 0; i < 3; i++ {
    fmt.Println(value, ":", i)
  }
}

// goroutine是一个轻量级的线程
func main () {
  // 普通调用
  test1("test1")
  // goroutine
  go test1("test2")

  // 创建一个匿名的goroutine
  go func (msg string) {
    fmt.Println(msg)
  }("test3")

  fmt.Scanln()
  fmt.Println("done")
}
