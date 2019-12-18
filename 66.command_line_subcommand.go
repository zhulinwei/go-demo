package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

//go run 66.command_line_subcommand.go flag1 -num=10 tony
func main ()  {
	flag1 := flag.NewFlagSet("flag1", flag.ExitOnError)
	flag1Num := flag1.Int("num", 10, "an int")

	flag2 := flag.NewFlagSet("flag2", flag.ExitOnError)

	if len(os.Args) < 2 {
		log.Fatalln("expected flag1 or flag2")
	}

	switch os.Args[1] {
	case "flag1":
		_ = flag1.Parse(os.Args[2:])
		fmt.Println("int:", *flag1Num)
		fmt.Println("all:", flag1.Args())
	case "flag2":
		_ = flag2.Parse(os.Args[2:])
		fmt.Println("all:", flag1.Args())
	default:
		log.Fatalln("expected flag1 or flag2")
	}
}