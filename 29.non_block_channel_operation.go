package main

import "fmt"

// 老实说，没太懂
func main () {
  message1 := make(chan string)
  message2 := make(chan string)
  select {
  case message3 := <- message1:
    fmt.Println("received message", message3)
  default: 
    fmt.Println("no message received")
  }

  message4 := "hello"
  select {
  case message1 <- message4:
    fmt.Println("send message", message1)
  default:
    fmt.Println("no message send")
  }

  select {
  case message5 := <- message1:
    fmt.Println("received message", message5)
  case message6 := <- message2:
    fmt.Println("received message", message6)
  default:
    fmt.Println("no activity")
  }
}
