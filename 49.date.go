package main

import "fmt"
import "time"

func main () {
  p := fmt.Println
  
  now := time.Now()
  p(now)

  then := time.Date(2010, 01, 01, 12, 30,30, 651387237, time.UTC)
  p(then)

  p(then.Year())
  p(then.Month())
  p(then.Day())
  p(then.Hour())
  p(then.Minute())
  p(then.Second())
  p(then.Nanosecond())
  // 当地时间的标准
  p(then.Location())

  p(then.Weekday())
  p(then.Before(now))
  p(then.After(now))
  p(then.Equal(now))

  // 计算两个时间差
  sub := now.Sub(then)
  p(sub)
  p(sub.Hours())
  p(sub.Minutes())
  p(sub.Seconds())
  p(sub.Nanoseconds())
  p(then.Add(sub))
  p(then.Add(-sub))
}
