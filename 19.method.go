package main

import "fmt"

// 我们定义一个矩形类型
type rect struct {
  width, height int
}

// 获取面积大小
func (rect rect) area() int {
  return rect.width * rect.height
}

func (rect rect) perim() int {
  return 2 * rect.width + 2 * rect.height
}

func main () {
  rect1 := rect{ width:10, height: 5 }
  fmt.Println("area:", rect1.area())
  fmt.Println("perim:", rect1.perim())

  rectPointer := &rect1
  fmt.Println("area:", rectPointer.area())
  fmt.Println("perim:", rectPointer.perim())
}
