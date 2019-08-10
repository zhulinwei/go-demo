package main

import "fmt"
import "sort"

func main () {
  // 字符串排序
  strings := []string{"c", "b", "a"}
  sort.Strings(strings)
  fmt.Println("strings:", strings)

  // 数字排序
  numbers := []int{3, 2, 1}
  sort.Ints(numbers)
  fmt.Println("numbers:", numbers)

  // 检查是否有序
  isSort := sort.IntsAreSorted(numbers)
  fmt.Println("sorted:", isSort)
}
