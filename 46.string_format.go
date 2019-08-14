package main

import "os"
import "fmt"

type point struct {
  x, y int
}

func main () {
  p := point{1, 2}

  fmt.Printf("%v\n", p)
  fmt.Printf("%+v\n", p)
  fmt.Printf("%#v\n", p)
  fmt.Printf("%T\n", p)
  fmt.Printf("%t\n", true)
  fmt.Printf("%d\n", 1234)
  fmt.Printf("%b\n", 1234)
  fmt.Printf("%c\n", 1234)
  fmt.Printf("%x\n", 1234)
  fmt.Printf("%f\n", 1234.0)
  fmt.Printf("%e\n", 1234.0)
  fmt.Printf("%E\n", 1234.0)
  fmt.Printf("%s\n", "\"string\"")
  fmt.Printf("%q\n", "\"string\"")
  fmt.Printf("%x\n", "hello world")
  fmt.Printf("%p\n", &p)
  fmt.Printf("|%6d|%6d|\n", 12, 345)
  fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)
  fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)
  fmt.Printf("|%6s|%6s|\n", "foo", "b")
  fmt.Printf("|%-6s|%-6s|\n", "foo", "b")

  s := fmt.Sprintf("a %s", "string")
  fmt.Println(s)
  fmt.Fprintf(os.Stderr, "an %s\n", "error")
}
