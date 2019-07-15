package main

import "fmt"
import "math"

// 我们定义一个数学几何体
type geometry interface {
  area() float64
  perim() float64
}

type rect struct {
  width, height float64
}

type circle struct {
  radius float64
}

func (rect rect) area() float64 {
  return rect.width * rect.height
}

func (rect rect) perim() float64 {
  return 2 * rect.width + 2 * rect.height
}

func (circle circle) area() float64 {
  return math.Pi * circle.radius * circle.radius
}

func (circle circle) perim() float64 {
  return 2 * math.Pi * circle.radius
}

func measure (geometry geometry) {
  fmt.Println(geometry)
  fmt.Println(geometry.area())
  fmt.Println(geometry.perim())
}

func main () {
  rect1 := rect{ width: 10, height: 10 }
  circle1 := circle{ radius: 10 }
  measure(rect1)
  measure(circle1)
}
