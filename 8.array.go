package main 

import "fmt"

func main () {
  var arr1 [5]int
  fmt.Println("array1:", arr1)

  arr1[4] = 100
  fmt.Println("set:", arr1)
  fmt.Println("get:", arr1[4])

  fmt.Println("length:", len(arr1))

  arr2 := [5]int{1, 2, 3, 4, 5}
  fmt.Println("array2:", arr2)

  var twoD [2][3]int
  for i := 0; i < 2; i++ {
    for j:= 0; j < 3; j++ {
      twoD[i][j] = i + j
    }
  }
  fmt.Println("2d:", twoD)
}
