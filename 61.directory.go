package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main () {
	// 创建文件夹
	if err := os.Mkdir("subdir", 0755); err != nil {
		log.Fatalf("mkdir fail: %v", err)
	}

	createEmptyFile := func (name string) {
		bytes := []byte("")
		if err := ioutil.WriteFile(name, bytes, 0644); err != nil {
			log.Fatalf("create empty file fail: %v", err)
		}
	}

	//创建文件
	createEmptyFile("subdir/file1")
	createEmptyFile("subdir/file2")
	createEmptyFile("subdir/file3")

	//创建文件，如果中间不存在对应的文件夹，则直接创建该文件夹
	if err := os.MkdirAll("subdir/child/file", 0755); err != nil {
		log.Fatalf("mkdir all file fail: %v", err)
	}

	// 读取指定文件夹下的文件/文件夹内容
	files, _ := ioutil.ReadDir("subdir")
	for _, file := range files {
		fmt.Println(" ", file.Name(), file.IsDir())
	}

	// Chdir相当于cd的作用
	if err := os.Chdir("subdir/child"); err != nil {
		log.Fatalf("chdir file fail: %v", err)
	}

	chileFiles, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatalf("chdir file fail: %v", err)
	}
	for _, file := range chileFiles {
		fmt.Println(" ", file.Name(), file.IsDir())
	}

	// 回到最上层以便后续删除文件夹失败
	_ = os.Chdir("../..")
	if err := os.RemoveAll("subdir"); err != nil {
		log.Fatalf("remove sub fail: %v", err)
	}
}