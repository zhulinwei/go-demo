package main 

import "os"
import "fmt"
import "strings"

func main () {
  _ = os.Setenv("FOO", "1")
  fmt.Println("FOO", os.Getenv("FOO"))
  fmt.Println("BOO", os.Getenv("BOO"))

  for _, env := range os.Environ() {
    pair := strings.Split(env, "=")
    fmt.Println(pair)
  }
}
