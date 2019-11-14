package main

import "fmt"
import "bufio"
import "net/http"

func main () {
  response, error := http.Get("http://www.baidu.com")
  if error != nil {
    panic(error)
  }
  defer response.Body.Close()
  fmt.Println("response status:", response.Status)
  scanner := bufio.NewScanner(response.Body)
  // 打印正文前5行
  for i := 0; scanner.Scan() && i < 5; i++ {
    fmt.Println(scanner.Text())
  }
  if error := scanner.Err(); error != nil {
    panic(error)
  }
}
