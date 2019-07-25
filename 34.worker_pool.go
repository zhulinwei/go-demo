package main

import "fmt"
import "time"

func worker (id int, jobs <- chan int, results chan <- int) {
  for i := range jobs {
    fmt.Println("worker", id, "started job", i)
    time.Sleep(time.Second)
    fmt.Println("worker", id, "finished job", i)
    results <- 2 * i
  }
}

// 模拟channel和gorunines交叉使用
func main () {
  jobs := make(chan int, 100)
  results := make(chan int, 100)

  for work := 1; work <= 3; work++ {
    go worker(work, jobs, results)  
  }

  for num := 1; num <= 5; num++ {
    jobs <- num 
  }
  close(jobs)

  for result := 1; result <= 5; result++ {
    <- results
  }
}
