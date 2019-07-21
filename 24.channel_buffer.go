package main

import "fmt"

func main () {
  // 初始化一个channel并设置可缓存2个值
  messages1 := make(chan string, 2)
  messages1 <- "test1"
  messages1 <- "test2"
  fmt.Println(<-messages1)
  fmt.Println(<-messages1)

  messages2 := make(chan int, 2)
  // 发送
  go func () {
    for i := 0; i < 10; i++ {
      fmt.Println("send:", i)
      messages2 <- i
    } 
    fmt.Println("close channel")
    close(messages2)
  }()

  // 接收
  for {
    elem, ok := <-messages2
    if !ok {
      fmt.Println("channel had closed")
      break
    }
    fmt.Println("receive:", elem)
  }
  fmt.Println("end!")
}


