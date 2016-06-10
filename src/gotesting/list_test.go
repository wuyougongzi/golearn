package gotesting

import (
	"fmt"
	"gostl"
	"testing"
)

//test function
func condition(v interface{}) bool {
	if v == "lbj" {
		return true
	}
	return false
}

func TestList(t *testing.T) {

	l := new(gostl.List)
	l.Insert(0, "apple")
	//l.Remove("apple")

	l.Erase(0)

	l.Push_Back("orange")
	l.Push_Front("bababa")

	//l.Remove("bababa")

	l.Insert(0, "peats")
	//l.Erase(2)

	fmt.Println("l.at: ", l.At(2))

	l.Insert(0, "lbj")
	l.Insert(0, "tmac")

	fmt.Println("..........next start............")

	node := l.Begin()
	for node != nil {
		fmt.Println(node.Value())
		node = node.Next()
	}

	fmt.Println("..........next end..............")

	fmt.Println(l.Size())

	/*	l.Clear()
		node = l.head
		for node != nil {
			fmt.Println(node.value)

			node = node.next

		}
	*/
}
