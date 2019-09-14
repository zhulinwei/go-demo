package main

import "fmt"
import "crypto/sha1"

func main () {
  strings := "sha1 this string"
  // 创建哈希实体
  hash := sha1.New()
  // Write期望得到的是字节
  hash.Write([]byte(strings))
  hashValue := hash.Sum(nil)
  fmt.Println(strings)
  // 打印16进制字符串
  fmt.Printf("%x\n", hashValue)
}
