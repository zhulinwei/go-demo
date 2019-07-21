package main

import "fmt"
import "time"
import "math/rand"

func test1 () {
  // 准备几个通道
  channels := [3]chan int {
    make(chan int, 1),
    make(chan int, 1),
    make(chan int, 1),
  }

  // 随机选择一个通道，并向他发送元素
  index := rand.Intn(3)
  fmt.Println("index is:", index)
  channels[index] <- index

  // 哪个通道有可取得元素值，哪个对应的分支就会被执行
  // 只要添加default分支，select块就不会被阻塞
  select {
  case <-channels[0]:
    fmt.Println("the first candidate case is selected.")
  case <-channels[1]:
    fmt.Println("the second candidate case is selected.")
  case <-channels[2]:
    fmt.Println("the third candidate case is selected.")
  default:
    fmt.Println("no candidate case is selected.")
  }
}

func test2 () {
  channel := make(chan int, 1)
  // 一秒后关闭通道
  time.AfterFunc(time.Second, func() {
    close(channel) 
  })

  select {
  case _, ok := <-channel:
    if !ok {
      fmt.Println("the candidate case is closed")
      break
    }
    fmt.Println("the candidate case is selected")
  }
}

// 添加默认分支后可以不用担心select被阻塞
func main () {
  test1()
  test2()
}
