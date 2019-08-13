package main

import "fmt"
import "strings"

// 搜索目标值索引
func Index(strings []string, target string) int {
  for index, value := range strings {
    if value == target {
      return index
    }
  }
  return -1
}

// 检查是否包含目标值
func Include (strings []string, target string) bool {
  return Index(strings, target) >= 0
}

// 检查是否存在符合目标函数的值
func Any (strings []string, f func (string) bool) bool {
  for _, value := range strings {
    if f(value) {
      return true
    }
  }
  return false
}

// 检查是否全部符合目标函数的值
func All (strings []string, f func (string) bool) bool {
  for _, value := range strings {
    if !f(value) {
      return false
    }
  }
  return true
}

// 过滤数组
func Filter (strings []string, f func (string) bool) []string {
  values := make([]string, 0)
  for _, value := range strings {
    if f(value) {
      values = append(values, value)
    }
  }
  return values
}

// 映射
func Map (strings []string, f func (string) string) []string {
  maps := make([]string, len(strings))

  for index, value := range strings {
    maps[index] = f(value)
  }
  return maps
}

// 一些常用的集合函数
func main () {
  var array = []string{"a", "b", "c", "d"}
  
  fmt.Println(Index(array, "a"))
  fmt.Println(Include(array, "a"))
  fmt.Println(Any(array, func (value string) bool {
    return strings.HasPrefix(value, "a")
  }))
  fmt.Println(All(array, func (value string) bool {
    return strings.HasPrefix(value, "a")
  }))
  fmt.Println(Filter(array, func (value string) bool {
    return strings.Contains(value, "a")
  }))
  fmt.Println(Map(array, strings.ToUpper))
}
