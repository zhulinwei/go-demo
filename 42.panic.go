package main

import "fmt"
import "errors"

// panic抛出错误，使用recover捕捉错误
func main () {
  defer func () {
    if err := recover(); err != nil {
      fmt.Printf("panic: %s\n", err)
    }
  }()
  panic(errors.New("just a problem"))
}
