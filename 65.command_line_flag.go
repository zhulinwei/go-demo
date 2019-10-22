package main

import (
	"flag"
	"fmt"
)

// go run 65.command_line_flag.go -num=12 -bool=true
func main () {
	numPtr := flag.Int("num", 10, "an int")
	boolPtr := flag.Bool("bool", false, "a bool")
	wordPtr := flag.String("word", "tony", "a string")

	flag.Parse()
	fmt.Println("num:", *numPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("word:", *wordPtr)
}