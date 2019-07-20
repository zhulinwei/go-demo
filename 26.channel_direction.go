package main

import "fmt"

// 将channel当做函数参数传递
func ping (pings chan <- string, msg string) {
  pings <- msg
}

func pong (pings <- chan string, pongs chan <- string) {
  message := <- pings
  pongs <- message
}

func main () {
  pings := make(chan string, 1)
  pongs := make(chan string, 1)
  ping(pings, "pass message")
  pong(pings, pongs)
  fmt.Println(<- pongs)
}
