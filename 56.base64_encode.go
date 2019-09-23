package main

import "fmt"
import b64 "encoding/base64"

func main () {
  data := "this is a string"

  // 标准编码器
  encode1 := b64.StdEncoding.EncodeToString([]byte(data))
  fmt.Println(encode1)
  decode1, _ := b64.StdEncoding.DecodeString(encode1)
  fmt.Println(string(decode1))

  // URL base64编码器
  encode2 := b64.URLEncoding.EncodeToString([]byte(data))
  fmt.Println(encode2)
  decode2, _ := b64.URLEncoding.DecodeString(encode2) 
  fmt.Println(string(decode2))
}
