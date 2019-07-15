package main

import "fmt"

// 参数可变型函数
func sum(nums ...int) {
  fmt.Print(nums, " ")
  total := 0
  for _, num := range nums {
    total += num
  }
  fmt.Println(total)
}

func main () {
  sum(1, 2)
  sum(1, 2, 3)

  nums := []int{1, 2, 3, 4}
  sum(nums...)
}
