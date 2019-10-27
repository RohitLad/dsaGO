package datastructures

import "fmt"

type LinkedList struct {
	head   *node
	tail   *node
	length int
}

type node struct {
	Data interface{}
	next *node
}

func (list *LinkedList) Append(elem interface{}) {
	newNode := node{Data: elem}
	if list.length == 0 {
		list.head = &newNode
		list.tail = list.head
	} else {
		lastNode := list.tail
		lastNode.next = &newNode
		list.tail = lastNode.next
	}
	list.length++
}

func (list *LinkedList) Pretend(elem interface{}) {
	newNode := node{Data: elem, next: list.head}
	list.head = &newNode
	list.length++
}

func (list *LinkedList) Insert(index int, elem interface{}) {
	if index == 0 {
		list.Pretend(elem)
	} else if index == list.length {
		list.Append(elem)
	} else {
		newNode := node{Data: elem}
		prevNode := list.goTo(index - 1)
		nextNode := prevNode.next
		prevNode.next = &newNode
		newNode.next = nextNode
		list.length++
	}

}

func (list *LinkedList) Remove(index int) {
	if index == 0 {
		list.head = list.head.next
	} else {
		prevNode := list.goTo(index - 1)
		node2Remove := prevNode.next

		if index == list.length-1 {
			list.tail = prevNode
			prevNode.next = nil
		} else {
			nextNode := node2Remove.next
			prevNode.next = nextNode
		}
	}
	list.length--
}

func (list LinkedList) goTo(index int) *node {
	currNode := list.head
	for i := 1; i <= index; i++ {
		currNode = currNode.next
	}
	return currNode
}

func (list LinkedList) Size() int {
	return list.length
}

func (list LinkedList) Display() {

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
