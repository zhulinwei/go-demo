package main

import (
	"fmt"
	"log"
	"path/filepath"
)

func main () {
	// Join拼接路径
	path1 := filepath.Join("demo", "go-demo", "60.file_path.go")
	fmt.Println(path1)
	//获取文件夹名称
	fmt.Println(filepath.Dir(path1))
	//获取文件名
	fmt.Println(filepath.Base(path1))
	//判断是否为绝对路径
	fmt.Println(filepath.IsAbs(path1))

	// 文件后缀名
	fmt.Println(filepath.Ext(path1))

	//获取目标地址的相对路径
	rel, err := filepath.Rel("a/c", "a/b/t/file")
	if err != nil {
		log.Fatalf("get rel fail:%v", err)
	}
	fmt.Println(rel)
}