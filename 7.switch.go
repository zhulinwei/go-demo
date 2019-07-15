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
