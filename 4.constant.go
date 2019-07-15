
package main

import "fmt"
import "math"

func main () {
  const test1 string = "string"
  fmt.Println(test1)

  const test2 = 3e20 / 5000000
  fmt.Println(test2)

  fmt.Println(int64(test2))

  fmt.Println(math.Sin(test2))
}
