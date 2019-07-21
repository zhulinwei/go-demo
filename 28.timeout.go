package main

import "fmt"
import "time"

func main () {
  channel1 := make(chan string, 1)
  go func () {
    time.Sleep(2 * time.Second)
    channel1 <- "test1"
  }()

  // timeout1会先比test1返回，因为sleep的时间较短
  select {
    case result1 := <- channel1:
      fmt.Println(result1)
    case <- time.After(1 * time.Second):
      fmt.Println("timeout 1")
  }

  // 注意：只要上面channel1结束后才会到channel2，即阻塞型
  channel2 := make(chan string, 1)
  go func () {
    time.Sleep(2 * time.Second)
    channel2 <- "test2"
  }()
  select {
    case result2 := <- channel2:
      fmt.Println(result2)
    case <- time.After(3 * time.Second):
      fmt.Println("timeout2")
  }
}
