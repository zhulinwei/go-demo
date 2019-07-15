package main

import "fmt"

func main () {
  strings1 := make([]string, 3)
  fmt.Println("strings: ", strings1)

  strings1[0] = "a"
  strings1[1] = "b"
  strings1[2] = "c"
  fmt.Println("set:", strings1)
  fmt.Println("get:", strings1)
  fmt.Println("length:", len(strings1))

  strings1 = append(strings1, "d")
  strings1 = append(strings1, "e", "f")
  fmt.Println("strings1 after append:", strings1)

  strings2 := make([]string, len(strings1))
  copy(strings2, strings1)
  fmt.Println("strings2 copy from strings1:", strings2)

  // 切割/复制动作，注意这是一个左闭右开的动作
  strings3 := strings1[2:5]
  fmt.Println("stings3:", strings3)

  // 复制下标0-4的元素
  strings4 := strings1[:5]
  fmt.Println("strings4:", strings4)

  // 复制下标2-最后的元素
  strings5 := strings1[2:]
  fmt.Println("strings5:", strings5)

  strings6 := []string{"g", "h", "i"}
  fmt.Println("strings6:", strings6)

  strings7 := make([][]int, 3)
  for i := 0; i < 3; i++ {
    innerLength := i + 1
    strings7[i] = make([]int, innerLength)
    for j := 0; j < innerLength; j++ {
      strings7[i][j] = i + j
    }
  }
  fmt.Println("2d:", strings7)
}
