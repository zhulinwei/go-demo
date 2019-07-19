package main

import "fmt"

func main () {
  // 初始化一个channel并设置可缓存2个值
  messages := make(chan string, 2)

  messages <- "test1"
  messages <- "test2"

  fmt.Println(<-messages)
  fmt.Println(<-messages)
}


