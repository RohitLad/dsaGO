package datastructures

import "fmt"

type Queue struct {
	data LinkedList
}

func (qu Queue) Size() int {
	return qu.data.Size()
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

func TestQueue() {
	queue := Queue{}
	queue.Enqueue(55)
	queue.Enqueue(65.2)
	fmt.Println("First element: ", queue.Peek())
	queue.Enqueue("test")
	queue.Display()
	fmt.Println("Removed element: ", queue.Dequeue())
	queue.Display()
}
