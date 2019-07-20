package main

import "fmt"
import "time"

// select允许同时操作多个channel
func main () {
  channel1 := make(chan string)
  channel2 := make(chan string)

  go func () {
    time.Sleep(1 * time.Second)
    channel1 <- "test1"
  }()
  go func () {
    time.Sleep(2 * time.Second)
    channel1 <- "test2"
  }()
  
  for i := 0; i < 2; i++ {
    select {
    case message1 := <- channel1:
      fmt.Println("received:", message1)
    case message2 := <- channel2:
      fmt.Println("received:", message2)
    }
  }
}
