package main

import "fmt"

// 自定义类型struct
type person struct {
  age int
  name string
}

func main () {
  // 简写
  fmt.Println(person{20, "Level.Z"})
  // 全写
  fmt.Println(person{ age: 20, name: "Level.Z" })
  // 少写
  fmt.Println(person{ name: "Level.Z" })

  struct1 := person{ age: 20, name: "Level.Z" }
  fmt.Println(struct1.age)
  fmt.Println(struct1.name)

  // 使用&产生指针结构
  structPoint := &struct1
  fmt.Println(structPoint.age)
}
