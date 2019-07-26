package main 

import "fmt"
import "sync"
import "time"

func worker (id int, work *sync.WaitGroup) {
  fmt.Printf("Worker %d starting\n", id)
  time.Sleep(time.Second)
  fmt.Printf("Worker %d done\n", id)
  work.Done()
}

// 模拟并发请求
func main () {
  var work sync.WaitGroup
  for i := 1; i <= 5; i++ {
    // 加入到WaitGroup中
    work.Add(1)
    fmt.Println(work)
    // 需要通过指针将WaitGroup传递给函数
    go worker(i, &work)
  }
  // 等待请求处理完成
  work.Wait()
}
