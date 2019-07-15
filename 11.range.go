package main

import "fmt"

// range关键字用于for循环中迭代数组、切片、通道或集合的元素
func main () {
  sum := 0
  nums := []int{1, 2, 3}
  // _在这里指的是占位符
  for _, num := range nums {
    sum += num
  }
  fmt.Println("sum:", sum)

  // 迭代数组
  for index, num := range nums {
    if num == 3 {
      fmt.Println("index: ", index)  
    }
  }

  map1 := map[string]string{"a": "apple", "b": "banana"}
  for key, value := range map1 {
    fmt.Printf("%s -> %s\n", key, value)
  }

  // 获取集合key值
  for key := range map1 {
    fmt.Println("key:", key)
  }

  for index, code := range "go" {
    fmt.Println(index, code)
  }
}
