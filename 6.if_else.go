
package main

import "fmt"

func main () {
  if 7 % 2 == 0 {
    fmt.Println("7 is even")
  } else {
    fmt.Println("7 is odd")
  }

  if 8 % 4 == 0 {
    fmt.Println("8 is divisible by 4")
  }

  if number1 := 9; number1 < 0 {
    fmt.Println(number1, "is negative")
  } else if number1 < 10 {
    fmt.Println(number1, "has 1 digit")
  } else {
    fmt.Println(number1, "has multiple digits")
  }
}
