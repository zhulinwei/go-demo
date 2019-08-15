package main

import "fmt"
import "regexp"

func main () {
  match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
  fmt.Println(match)

  // 创建一个正则表达式
  regex, _ := regexp.Compile("p([a-z]+)ch")

  fmt.Println(regex.MatchString("peach"))
  fmt.Println(regex.FindString("peach punch"))
  fmt.Println(regex.FindStringIndex("peach punch"))
  fmt.Println(regex.FindStringSubmatch("peach punch"))
  fmt.Println(regex.FindStringSubmatchIndex("peach punch"))
  fmt.Println(regex.FindAllString("peach punch pinch", -1))
  fmt.Println(regex.FindAllStringSubmatchIndex("peach punch pinch", -1))
  fmt.Println(regex.FindAllString("peach punch pinch", 2))
  fmt.Println(regex.Match([]byte("peach")))
}
