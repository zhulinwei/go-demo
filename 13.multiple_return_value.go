package main

import "fmt"

// 抛出多值
func getValues ()(int, int) {
  return 3, 7 
}

func main () {
  // 获取多值
  value1, value2 := getValues()
  fmt.Println(value1)
  fmt.Println(value2)

  // 使用_忽略不想要的值
  _, value3 := getValues()
  fmt.Println(value3)
}
