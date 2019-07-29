package main

import "fmt"
import "sync"
import "time"
import "math/rand"
import "sync/atomic"

// 互斥锁
func main () {
  var state = make(map[int]int)
  var mutex = &sync.Mutex{}

  var read uint64
  var write uint64
  for i := 0; i < 100; i++ {
    go func () {
      total := 0
      for {
        key := rand.Intn(5)
        mutex.Lock()
        total += state[key]
        fmt.Println(total)
        mutex.Unlock()
        atomic.AddUint64(&read, 1)
        time.Sleep(time.Millisecond)
      }
    }()
  }

  for work := 0; work < 10; work++ {
    go func () {
      for {
        key := rand.Intn(5)
        value := rand.Intn(100)

        mutex.Lock()
        state[key] = value
        mutex.Unlock()
        atomic.AddUint64(&write, 1)
        time.Sleep(time.Millisecond)
      }
    }()
  }

  time.Sleep(time.Second)

  readResult := atomic.LoadUint64(&read)
  fmt.Println("read result:", readResult)
  writeResult := atomic.LoadUint64(&write)
  fmt.Println("write result:", writeResult)

  mutex.Lock()
  fmt.Println("state:", state)
  mutex.Unlock()
}
