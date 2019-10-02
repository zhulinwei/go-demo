package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// 直接写
	bytes := []byte("hello world")
	err := ioutil.WriteFile("./test1.txt", bytes, 0644)
	if err != nil {
		panic(err.Error())
	}

	// 追加写
	file, err := os.Create("./test2.txt")
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatalf("os create file fail: %v", err)
		}
	}()

	bytes2 := []byte{115, 111, 109, 101, 10}
	count2, err := file.Write(bytes2)
	fmt.Printf("wrote %d bytes\n", count2)

	count3, err := file.WriteString("writes\n")
	fmt.Printf("wrote %d bytes\n", count3)

	// sync将写入刷新到稳定的存储中
	if err := file.Sync(); err != nil {
		log.Printf("file sync fail: %v", err)
	}

	write := bufio.NewWriter(file)

	count4, err := write.WriteString("buffered\n")
	fmt.Printf("wrote %d bytes\n", count4)

	// flush确保所有操作缓冲到底层
	if err := write.Flush(); err != nil {
		log.Printf("write flush fail: %v", err)
	}
}
