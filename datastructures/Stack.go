package datastructures

import "fmt"

type Stack struct {
	data LinkedList
}

func (st *Stack) Push(elem interface{}) {
	st.data.Pretend(elem)
}

func (st *Stack) Peek() interface{} {
	return st.data.head.Data
}

func (st *Stack) Pop() interface{} {

	headNode := st.data.head
	st.data.Remove(0)
	return headNode.Data
}

func (st Stack) Display() {
	st.data.Display()
}

func TestStack() {
	stack := Stack{}
	stack.Push(55)
	stack.Push(65.2)
	fmt.Println("Top element: ", stack.Peek())
	stack.Push("test")
	stack.Display()
	fmt.Println("Popped element: ", stack.Pop())
	stack.Display()
}
