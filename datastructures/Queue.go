package datastructures

type Queue struct {
	data LinkedList
}

func (qu *Queue) Enqueue(elem interface{}) {
	qu.data.Append(elem)
}

func (qu *Queue) Dequeue() interface{} {
	lastNode := qu.data.head
	qu.data.Remove(0)
	return lastNode.Data
}

func (qu Queue) Peek() interface{} {
	return qu.data.head.Data
}

func (qu Queue) Display() {
	qu.data.Display()
}
