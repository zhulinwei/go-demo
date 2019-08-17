package main

import "fmt"
import "time"

// unix
func main () {
  now := time.Now()
  seconds := now.Unix()
  nanoseconds := now.UnixNano()
  millis := nanoseconds / 1000000
  fmt.Println(now)
  fmt.Println(seconds)
  fmt.Println(millis)
  fmt.Println(nanoseconds)

  fmt.Println(time.Unix(seconds, 0))
  fmt.Println(time.Unix(0, nanoseconds))
}
