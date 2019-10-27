package datastructures

type Stack struct {
	data DoubleLinkedList
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
