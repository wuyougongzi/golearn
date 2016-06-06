package gotesting

import (
	"fmt"
	"gostl"
	"testing"
)

func TestStack(t *testing.T) {
	s := new(gostl.Stack)
	fmt.Println("stack test")
	s.Push(5)
	fmt.Println("stack size: ", s.Size())
	fmt.Println("stack top: ", s.Top())

	s.Push(1)
	fmt.Println("stack size: ", s.Size())
	fmt.Println("stack top: ", s.Top())

	s.Push(3)
	fmt.Println("stack size: ", s.Size())
	fmt.Println("stack top: ", s.Top())

	s.Pop()
	fmt.Println("stack size: ", s.Size())
	fmt.Println("stack top: ", s.Top())

	s.Pop()
	fmt.Println(s.Empty())
	s.Pop()
	fmt.Println(s.Empty())

}
