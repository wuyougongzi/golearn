/*
go queue
*/
package gostl

type Queue struct {
	tail *Element
	head *Element
	size int
}

//be defined in go stack
/*type Element struct {
	next  *Element
	value interface{}
}
*/
func (q *Queue) Push(v interface{}) {
	e := new(Element)
	e.value = v
	if q.tail == nil {
		q.head = e
		q.tail = e
		q.head.next = q.tail
	} else {
		q.tail.next = e
		q.tail = e
	}

	q.size++
}

func (q *Queue) Pop() {
	if q.Empty() {
		panic("queue is empty")
	}

	//size 1 the next code is always effect
	if q.head != nil && q.head == q.tail {
		q.head.next = nil
	}
	q.head = q.head.next

	q.size--
}

func (q *Queue) Front() interface{} {
	if q.Empty() {
		panic("queue is empty")
	}

	return q.head.value
}

func (q *Queue) Empty() bool {
	return q.head == nil || q.size == 0
}

func (q *Queue) Size() int {
	return q.size
}
