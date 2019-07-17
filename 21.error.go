package main

import "fmt"
import "errors"

// 使用自带的错误
func test1 (arg int) (int, error) {
  if arg == 1 {
    // error 放在错误的最后一个位置上
    return -1, errors.New("when arg is 1, I will throw a error")
  }
  // nil 在出现在error的位置上，即表示没有错误
  return arg + 1, nil
}

type argError struct {
  arg int
  prob string
}

func (error *argError) Error() string {
  return fmt.Sprintf("%d - %s", error.arg, error.prob)
}

// 使用自定义错误
func test2 (arg int) (int, error) {
  if arg == 1 {
    return -1, &argError{arg, "when arg is 1, I will throw anthor error"}
  }
  return arg + 2, nil
}

func main () {
  for _, value := range []int{1,2} {
    if arg, error := test1(value); error != nil {
      fmt.Println("test1 failed:", error)
    } else {
      fmt.Println("test1 work:", arg)
    }
  }
 
  for _, value := range []int{1, 2} {
    if arg, error := test2(value); error != nil {
      fmt.Println("test2 failed:", error)
    } else {
      fmt.Println("test2 work:", arg)
    }
  }

  _, error := test2(1)
  if value, errorMsg := error.(*argError); errorMsg {
    fmt.Print(value.arg, value.prob)
  }
}
