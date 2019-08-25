package main

import "fmt"
import "net"
import "net/url"

func main () {
  path := "postgres://user:pass@host.com:5432/path?k1=v1&k2=v2#f"

  value, _ := url.Parse(path)
  fmt.Println(value.Scheme)
  fmt.Println(value.User)
  fmt.Println(value.User.Username())
  fmt.Println(value.User.Password())

  
  fmt.Println(value.Host)
  host, port, _ := net.SplitHostPort(value.Host)
  fmt.Println(host)
  fmt.Println(port)

  
  fmt.Println(value.Path)
  fmt.Println(value.Fragment)
  fmt.Println(value.RawQuery)

  maps, _ := url.ParseQuery(value.RawQuery)
  fmt.Println(maps)
  fmt.Println(maps["k1"])
  fmt.Println(maps["k1"][0])
}
