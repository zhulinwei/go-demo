package main

import "fmt"
import "time"
import "math/rand"
import "sync/atomic"

type readOption struct {
  key int
  response chan int
}

type writeOption struct {
  key int
  value int
  response chan bool
}

// 老实说，没完全懂，就是通过channel和goroutine实现与互斥锁一致的效果
func main () {
  var readOptions uint64
  var writeOptions uint64

  reads := make(chan readOption)
  writes := make(chan writeOption)

  go func () {
    var state = make(map[int]int)
    for {
      select {
      case read := <- reads:
        read.response <- state[read.key]
      case write := <- writes:
        state[write.key] = write.value
        write.response <- true
      }
    }
  }()

  for i := 0; i < 100; i++ {
    go func () {
      for {
        read := readOption{
          key: rand.Intn(5),
          response: make(chan int)}
        reads <- read
        <- read.response
        atomic.AddUint64(&readOptions, 1)
        time.Sleep(time.Millisecond)
      }
    }()
  }

  for j := 0; j < 10; j++ {
    go func () {
      for {
        write := writeOption{
          key: rand.Intn(5),
          value: rand.Intn(100),
          response: make(chan bool)}
        writes <- write
        <- write.response
        atomic.AddUint64(&writeOptions, 1)
      }
    }()
  }

  time.Sleep(time.Second)
  readResult := atomic.LoadUint64(&readOptions)
  fmt.Println("read result:", readResult)
  writeResult := atomic.LoadUint64(&writeOptions)
  fmt.Println("write result:", writeResult)
}
