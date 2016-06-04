package main

import (
	"fmt"
)

func main() {
	//define array
	var myArray [10]int = [10]int{1, 2, 3, 4, 10, 5, 6, 7, 8, 9}

	//create slice base on array
	var myslice []int = myArray[:5]

	fmt.Println("Element of myArray")
	for _, v := range myArray {
		fmt.Print(v, " ")
	}

	fmt.Println("Elemnt of mySlice")
	for _, v := range myslice {
		fmt.Print(v, " ")
	}
}
