/*
gostl list
two-way link lists
*/
package gostl

type Node struct {
	next  *Node
	prev  *Node
	value interface{}
}

//use for visit node
func (node *Node) Next() *Node {
	return node.next
}

func (node *Node) Value() interface{} {
	return node.value
}

type List struct {
	head *Node
	tail *Node
	size int
}

//add value at tail
func (l *List) Push_Back(value interface{}) {
	e := new(Node)
	e.value = value
	if l.head == nil {
		l.head = e
		l.head.next = l.tail
	} else {
		if l.tail == nil {
			l.tail = e
			l.tail.prev = l.head
			l.head.next = l.tail
		} else {
			l.tail.next = e
			e.prev = l.tail
		}
	}

	l.size++
}

//delete  last Node at tail
func (l *List) Pop_Back() {
	if l.IsEmpty() {
		panic("list is empty")
	}

	if l.tail == nil {
		l.head = nil
	} else {
		l.tail = l.tail.prev
		l.tail.next = nil
	}

	l.size--
}

//add Node at front
func (l *List) Push_Front(value interface{}) {
	e := new(Node)
	e.value = value
	e.next = l.head
	if l.head == nil {
		l.head = e
		l.head.next = l.tail
	} else {
		l.head.prev = e
		l.head = e
	}

	l.size++
}

//delete Node at front
func (l *List) Pop_Front() {
	if l.IsEmpty() {
		panic("list is empty")
	}

	if l.head == nil {
		panic("list is empty")
	} else {
		l.head = l.head.next
	}
	l.size--
}

//insert a elment behind pos
func (l *List) Insert(pos int, value interface{}) {
	if pos > l.size {
		panic("out of list range")
	}

	e := new(Node)
	e.value = value

	if l.head == nil {
		l.head = e
		l.head.next = l.tail
	} else {
		i := pos
		tmpElem := l.head
		for i > 0 && tmpElem.next != nil {
			tmpElem = tmpElem.next
			i--
		}

		//link
		e.prev = tmpElem.prev
		e.next = tmpElem
		if tmpElem.prev != nil {
			tmpElem.prev.next = e
		}
		tmpElem.prev = e

		if pos == 0 {
			l.head = e
		}
	}

	l.size++
}

//remove all Nodes which's value == value
func (l *List) Remove(value interface{}) {
	if l.head == nil {
		panic("list is empty")
	}

	tmpNode := l.head
	for tmpNode != nil {
		if tmpNode.value == value {
			if tmpNode.next != nil {
				tmpNode.next.prev = tmpNode.prev
			}

			if tmpNode.prev != nil {
				tmpNode.prev.next = tmpNode.next
			}

			if tmpNode == l.head {
				l.head = tmpNode.next
			}

			l.size--
		}
		tmpNode = tmpNode.next
	}
}

//remove all Nodes which's value makes func f true
func (l *List) Remove_If(f func(v interface{}) bool) {
	if l.head == nil {
		panic("list is empty")
	}

	tmpNode := l.head
	for tmpNode != nil {
		if f(tmpNode.value) {
			if tmpNode.next != nil {
				tmpNode.next.prev = tmpNode.prev
			}

			if tmpNode.prev != nil {
				tmpNode.prev.next = tmpNode.next
			}

			if tmpNode == l.head {
				l.head = tmpNode.next
			}

			l.size--
		}
		tmpNode = tmpNode.next
	}
}

//delete Node at pos
func (l *List) Erase(pos int) {
	if l.head == nil {
		panic("list is empty")
	}

	if pos >= l.size {
		panic("out of list range")
	}

	tmpNode := l.head
	for pos > 0 && tmpNode.next != nil {
		tmpNode = tmpNode.next
		pos--
	}

	if tmpNode.next != nil {
		tmpNode.next.prev = tmpNode.prev
	}
	if tmpNode.prev != nil {
		tmpNode.prev.next = tmpNode.next
	}

	if tmpNode == l.head {
		l.head = tmpNode.next
	}

	l.size--
}

func (l *List) Clear() {
	l.head = nil
	l.tail = nil
	l.size = 0
}

//
func (l *List) IsEmpty() bool {
	if l.head == nil || l.size == 0 {
		return true
	}
	return false
}

func (l *List) Size() int {
	return l.size
}

func (l *List) At(pos int) interface{} {
	if l.head == nil {
		panic("list is empty")
	}
	if pos >= l.size {
		panic("out of list range")
	}

	tmpNode := l.head
	for pos > 0 && tmpNode != nil {
		tmpNode = tmpNode.next
		pos--
	}
	return tmpNode.value
}

func (l *List) Begin() *Node {
	return l.head
}
