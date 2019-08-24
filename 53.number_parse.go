package main

import "fmt"
import "strconv"

// 内置宝strconv提供数字解析
func main () {
  num1, _ := strconv.ParseFloat("1.234", 64)
  fmt.Println(num1)

  num2, _ := strconv.ParseInt("123", 0, 64)
  fmt.Println(num2)

  num3, _ := strconv.ParseInt("0x1c8", 0, 64)
  fmt.Println(num3)

  num4, _ := strconv.ParseUint("789", 0, 64)
  fmt.Println(num4)

  num5, _ := strconv.Atoi("123")
  fmt.Println(num5)
}
