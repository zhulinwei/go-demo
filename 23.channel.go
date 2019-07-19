package main

import "fmt"

// channel用来在goroutine中传递数据，相当于管道的作用
func main () {
  // 创建一个channel
  message1 := make(chan string)
  fmt.Println(message1)
  
  // 往channel中塞东西
  go func () { message1 <- "message" }()

  // 从channel中取东西
  message2 := <-message1
  fmt.Println(message2)

  // 再取，取不到啦，因为goroutine已经sleep了，此时channel被锁住啦
  // message3 := <-message1
  // fmt.Println(message3)
}
