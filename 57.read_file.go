package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main () {
	data, err := ioutil.ReadFile("./1.hello_world.go")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(string(data))

	file, err := os.Open("./1.hello_world.go")
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(file.Name())

	byte1 := make([]byte, 25)
	n1, err := file.Read(byte1)
	fmt.Println(string(byte1[:n1]))
}