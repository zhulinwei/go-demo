package main

import "fmt"
import "time"

func main () {
  number1 := 2
  fmt.Print("Write ", number1, " as ")
  switch number1 {
    case 1: 
      fmt.Println("one")
    case 2: 
      fmt.Println("two")
    case 3:
      fmt.Println("three")
  }
  
  // fallthrough会继续执行下一条case,不会再判断下一条case的expr,如果之后的case都有fallthrough,default出会被执行
  // *fallthrough只能在句尾，不能在中间
  switch  {
    case number1 < 2: 
      fmt.Println("one")
      fallthrough
    case number1 = 2: 
      fmt.Println("two")
    case 3:
      fmt.Println("three")
  }
  
  switch time.Now().Weekday() {
    case time.Saturday, time.Sunday: 
      fmt.Println("It's the weekend")
    default:
      fmt.Println("It's a weekday")
  }

  nowTime := time.Now()
  switch {
    case nowTime.Hour() < 12:
      fmt.Println("It's before noon")
    default: 
      fmt.Println("It's after noon")
  }

}
