package testing

import (
	"fmt"
	"gostl"
	"testing"
)

func TestQueue(t *testing.T) {
	q := new(gostl.Queue)
	//	fmt.Println("queue front: ", q.Front())
	q.Push(1)
	//	q.Push(2)
	fmt.Println("queue front: ", q.Front())

	q.Push(5)
	fmt.Println("queue front: ", q.Front())
	fmt.Println("queue size: ", q.Size())

	q.Pop()
	fmt.Println("queue front: ", q.Front())
	fmt.Println("queue size: ", q.Size())

	q.Push(2)
	fmt.Println("queue front: ", q.Front())

	fmt.Println("queue Empty: ", q.Empty())
	q.Pop()
	fmt.Println("queue Empty: ", q.Empty())
	q.Pop()
	fmt.Println("queue Empty: ", q.Empty())
	//	exit(0)
}
