package datastructures

import "fmt"

type DoubleLinkedList struct {
	head   *DNode
	tail   *DNode
	length int
}

type DNode struct {
	Data interface{}
	next *DNode
	prev *DNode
}

func (list *DoubleLinkedList) Append(elem interface{}) {
	newNode := DNode{Data: elem}
	if list.length == 0 {
		list.head = &newNode
		list.tail = &newNode
	} else {
		lastNode := list.tail
		lastNode.next = &newNode
		newNode.prev = lastNode
		list.tail = &newNode
	}
	list.length++
}

func (list *DoubleLinkedList) Pretend(elem interface{}) {
	newNode := DNode{Data: elem, next: list.head}
	if list.length > 0 {
		list.head.prev = &newNode
	}

	list.head = &newNode
	list.length++
}

func (list *DoubleLinkedList) Insert(index int, elem interface{}) {
	if index == 0 {
		list.Append(elem)
	} else if index == list.length {
		list.Append(elem)
	} else {
		newNode := DNode{Data: elem}
		prevNode := list.goTo(index - 1)
		nextNode := prevNode.next
		prevNode.next = &newNode
		newNode.next = nextNode
		nextNode.prev = &newNode
		newNode.prev = prevNode
		list.length++
	}
}

func (list *DoubleLinkedList) Remove(index int) {

	if index == 0 {
		nextNode := list.head.next
		nextNode.prev = nil
		list.head = nextNode
	} else {
		node2Remove := list.goTo(index)
		prevNode := node2Remove.prev

		if index == list.length-1 {
			list.tail = prevNode
			prevNode.next = nil
		} else {
			nextNode := node2Remove.next
			nextNode.prev = prevNode
			prevNode.next = nextNode
		}
	}
	list.length--
}

func (list *DoubleLinkedList) Reverse() {
	if list.length > 1 {
		currNode := list.head
		for i := 0; i <= list.length-1; i++ {
			tempNode := currNode.next
			currNode.next = currNode.prev
			currNode.prev = tempNode
			currNode = tempNode
		}
		tmpNode := list.head
		list.head = list.tail
		list.tail = tmpNode
	}
}

func (list DoubleLinkedList) goTo(index int) *DNode {
	if index <= list.length/2 {
		currNode := list.head
		for i := 1; i <= index; i++ {
			currNode = currNode.next
		}
		return currNode
	} else {
		currNode := list.tail
		for i := list.length - 2; i >= index; i-- {
			currNode = currNode.prev
		}
		return currNode
	}
}

func (list DoubleLinkedList) Size() int {
	return list.length
}

func (list DoubleLinkedList) Display() {
	node := list.head
	if node == nil {
		fmt.Println("[]")
	} else {
		for node != nil {
			fmt.Printf("%v ->", node.Data)
			node = node.next
		}

	}
	fmt.Println("   Size: ", list.Size())
}

func TestDoubleLinkedList() {
	list := DoubleLinkedList{}
	list.Append(55.0)
	list.Append(88)
	list.Append(68.125)
	list.Pretend(458)
	list.Display()
	list.Reverse()
	list.Display()
	list.Insert(2, 33.5)
	list.Display()
	list.Remove(4)
	list.Append(78)
	list.Remove(1)
	list.Display()
	list.Reverse()
	list.Display()
}
