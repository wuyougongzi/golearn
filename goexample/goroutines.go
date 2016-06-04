//A goroutines is a lightweight thread of excution

package main

import (
	"fmt"
)

func fLikeThread(from string) {
	fmt.Println(from, " start")
	for i := 0; i < 10; i++ {
		fmt.Println(from, ": ", i)
	}
}

func main() {
	go fLikeThread("direct")

	go fLikeThread("goroutines")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	var input string
	fmt.Scanln(&input)
	fmt.Println("done")

}
