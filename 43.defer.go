package main

import "fmt"

// defer可以延迟函数的执行
// 多个defer函数被定义时，函数的调用的顺序和定义的顺序相反
func main() {
	defer fmt.Println("first defer")
	for i := 0; i < 3; i++ {
		defer fmt.Printf("defer in for [%d]\n", i)
	}
	defer fmt.Println("last defer")
}
