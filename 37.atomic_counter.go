package main 

import "fmt"
import "time"
import "sync/atomic"

// 原子计数器
func main () {
  var count uint64
  for i := 0; i < 50; i++ {
    go func () {
      // 使用原子方式递增计数器，并使用&语法为其提供计数器的内存地址
      atomic.AddUint64(&count, 1)
      fmt.Println(count)
      time.Sleep(time.Millisecond)
    }()  
  }
  time.Sleep(time.Second)

  countResult := atomic.LoadUint64(&count)
  fmt.Println("count:", countResult)
}
