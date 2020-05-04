package main

import "os"
import "fmt"

// 注意：defer不会被执行
func main () {
  defer fmt.Println("defer")

  os.Exit(3)
}
