package main

import (
	"fmt"
	"os"
)

func main() {

	length := len(os.Args)
	fmt.Println(length)

	for i, arg := range os.Args {
		fmt.Println(i, ": ", arg)
	}

	argWithProg := os.Args
	argWithoutProg := os.Args[1:]

	arg := os.Args[3]

	fmt.Println(argWithProg)
	fmt.Println(argWithoutProg)
	fmt.Println(arg)
}
