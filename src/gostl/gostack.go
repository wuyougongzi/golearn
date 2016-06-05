/*
go stack struct
*/
package gostl

type Stack struct {
	top  *Element
	size int
}

type Element struct {
	next  *Element
	value interface{}
}

func (s *Stack) Push(v interface{}) {
	e := new(Element)
	e.value = v
	e.next = s.top
	s.top = e
	s.size++
}

func (s *Stack) Pop() {
	if s.top == nil || s.size == 0 {
		panic("stack is empty, can not pop")
	}

	newTop := s.top.next
	s.top = newTop
	s.size--
}

func (s Stack) Top() interface{} {
	if s.top == nil {
		panic("top is nil error")
	}

	return s.top.value
}

func (s Stack) Empty() bool {
	return s.size == 0
}

func (s Stack) Size() int {
	return s.size
}
