package main

import (
	"fmt"
)

func main() {
	str := "Hello,世界"
	fmt.Println(str)

	for i, ch := range str {
		fmt.Println(i, ch)
	}
	fmt.Println(".......................")
	for i := 0; i < len(str); i++ {
		ch := str[i]
		fmt.Println(i, ch)
	}
}
