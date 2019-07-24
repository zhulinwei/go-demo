package main

import "fmt"
import "time"

func main () {
  ticker := time.NewTicker(500 * time.Millisecond)

  go func () {
    for time1 := range ticker.C {
      fmt.Println("Ticker at", time1)
    }
  }()
  
  time.Sleep(1600 * time.Millisecond)
  ticker.Stop()
  fmt.Println("Ticker stopped")
}
