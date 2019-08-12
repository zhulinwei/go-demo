package main

import "fmt"
import "sort"

type byLength []string

func (value byLength) Len() int {
  return len(value)
}

func (value byLength) Swap(i, j int) {
  value[i], value[j] = value[j], value[i]
}

func (value byLength) Less(i, j int) bool {
  return len(value[i]) > len(value[j])
}

func main () {
  fruits := []string{"peach", "banana", "kiwi"}
  sort.Sort(byLength(fruits))
  fmt.Println(fruits)
}
