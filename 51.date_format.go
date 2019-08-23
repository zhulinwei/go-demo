package main

import "fmt"
import "time"

func main () {
  p := fmt.Println

  t := time.Now()
  p(t.Format(time.RFC3339))
  time1, _ := time.Parse(
   time.RFC3339,
   "2012-11-01T22:08:41+00:00")
  p(time1)

  p(t.Format("3:04PM"))
  p(t.Format("Mon Jan _2 15:04:05 2006"))
  p(t.Format("2006-01-02T15:04:05.999999-07:00"))

  form := "3 04 PM"
  time2, _ := time.Parse(form, "8 41 PM")
  p(time2)

  fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
    t.Year(), t.Month(), t.Day(),
    t.Hour(), t.Minute(), t.Second())

  ansic := "Mon Jan _2 15:04:05 2006"
  time3, _ := time.Parse(ansic, "8:41PM")
  p(time3)
}
