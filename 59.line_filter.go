package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// 读取os的输入，我们可以尝试ls | go run 59.line_filter.go命令启动项目
func main () {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println(scanner.Text())

	for scanner.Scan() {
		ucl := strings.ToUpper(scanner.Text())

		fmt.Println(ucl)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error: %v", err)
	}
}