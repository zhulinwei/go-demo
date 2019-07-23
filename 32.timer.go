package main

import "fmt"
import "time"

func main () {
  // 定义一个2秒的定时器,延迟执行
  timer1 := time.NewTimer(2 * time.Second)

  // 执行等待的过程
  <- timer1.C
  fmt.Println("Timer 1 expired")

  timer2 := time.NewTimer(time.Second)
  go func () {
    <- timer2.C
    fmt.Println("Timer 2 expired")
  }()

  // stop后拿到的是一个布尔值
  stop2 := timer2.Stop()
  if stop2 {
    fmt.Println("Timer 2 stopped")
  }
}
