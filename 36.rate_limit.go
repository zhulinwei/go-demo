package main

import "fmt"
import "time"

// 并发限制
func main () {
  requests1 := make(chan int, 5)
  for i := 1; i <= 5; i++ {
    requests1 <- i
  }
  close(requests1)

  limiter1 := time.Tick(200 * time.Millisecond)

  for request1 := range requests1 {
    // 循环时我们每次阻塞200ms
    <-limiter1
    fmt.Println("request", request1, time.Now())
  }

  limiter2 := make(chan time.Time, 3)
  for i := 0; i < 3; i++ {
    limiter2 <- time.Now()
  }
  go func () {
    for tick := range time.Tick(200 * time.Millisecond) {
      limiter2 <- tick
    }
  }()

  requests2 := make(chan int, 5)
  for i := 1; i <= 5; i++ {
    requests2 <- i 
  }
  close(requests2)
  for request2 := range requests2 {
    <- limiter2
    fmt.Println("request", request2, time.Now())
  }
}
