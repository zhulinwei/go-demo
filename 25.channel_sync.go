package main 

import "fmt"
import "time"

// done通道用于通知另一个goroutine功过已完成
func worker (done chan bool) {
  fmt.Print("working...")
  time.Sleep(time.Second)
  fmt.Println("done")

  done <- true
}

func main () {
  done := make(chan bool, 1)
  go worker(done)
  
  <- done
}
