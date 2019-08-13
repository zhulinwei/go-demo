package main

import "fmt"
import s "strings"

var p = fmt.Println

// 一些常见的字符串函数
func main () {
  p("Contains:", s.Contains("test", "t"))
  p("Count:", s.Count("test", "t"))
  p("HasPrefix:", s.HasPrefix("test", "te"))
  p("HasSuffix:", s.HasSuffix("test", "st"))
  p("Index:", s.Index("test", "e"))
  p("Join:", s.Join([]string{"a", "b", "c"}, "-"))
  p("Repeat:", s.Repeat("a", 5))
  p("Replace:", s.Replace("foo", "o", "0", 1))
  p("Replace:", s.Replace("foo", "o", "0", -1))
  p("Split:", s.Split("a-b-c-d-e", "-"))
  p("ToLower:", s.ToLower("TEST"))
  p("ToUpper:", s.ToUpper("test"))
  p("Len:", len("hello"))
  p("Char:", "hello"[2])
}
