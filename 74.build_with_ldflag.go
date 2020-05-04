package main

import "fmt"

var Foo string

// go build -ldflags "-X 'main.Foo=test'" 74.build_with_ldflag.go
func main() {
	if Foo == "" {
		fmt.Println("Foo is empty.")
	} else {
		fmt.Printf("Foo=%s\n", Foo)
	}
}
