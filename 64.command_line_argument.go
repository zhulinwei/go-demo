package main

import (
	"fmt"
	"os"
)

// 读取启动参数
// go run 64.command_line_argument.go a b c d
func main () {
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	arg := os.Args[3]
	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}