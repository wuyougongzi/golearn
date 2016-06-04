package main

import (
	"fmt"
)

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perm(val int) int {
	return r.height*val + r.width*2
}

func main() {
	r := rect{width: 10, height: 5}
	fmt.Println(r.area())

	fmt.Println(r.perm(2))
}
