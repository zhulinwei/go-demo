package main 

import "fmt"

func plus1 (value1 int, value2 int) int {
  return value1 + value2 
}

// 多参数类型相同时，可省略前面的类型
func plus2 (value1, value2, value3 int) int {
  return value1 + value2 + value3
} 

func main () {
  result1 := plus1(1, 2)
  fmt.Println("1+2 =", result1)

j result2 := plus2(1, 2, 3)
  fmt.Println("1+2+3 =", result2)
}
