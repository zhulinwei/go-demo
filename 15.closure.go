
package main 

import "fmt"

func intSeq() func() int {
  num := 0
  return func () int {
    num++
    return num
  }
}

func main () {
  nextInt := intSeq()

  // num开始累加
  fmt.Println(nextInt())
  fmt.Println(nextInt())
  fmt.Println(nextInt())
  fmt.Println(nextInt())
  fmt.Println(nextInt())

  // 新的闭包，num重新开始累加
  newInts := intSeq()
  fmt.Println(newInts())
}
