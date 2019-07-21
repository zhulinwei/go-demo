package main

import "fmt"

func ping (pings chan<- string, msg string) {
  pings <- msg
}

// pings <-cnan string：约束只能发
// pongs chan<- string: 约束只能收
func pong (pings <-chan string, pongs chan<- string) {
  message := <-pings
  pongs <- message
}

// channel的约束（收、发、收/发）
func main () {
  pings := make(chan string, 1)
  pongs := make(chan string, 1)
  ping(pings, "pass message")
  pong(pings, pongs)
  fmt.Println(<- pongs)
}
