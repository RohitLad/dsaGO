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
		currNode := list.head
		for i := 1; i <= index-1; i++ {
			currNode = currNode.next
		}
		newNode.next = currNode.next
		currNode.next = &newNode
		list.length++
	}

}

func (list *LinkedList) Remove(index int) {
	if index == 0 {
		list.head = list.head.next
	} else {
		currNode := list.head
		for i := 1; i <= index-1; i++ {
			currNode = currNode.next
		}
		removeNode := currNode.next
		next2removeNode := removeNode.next
		currNode.next = next2removeNode
		if index == list.length-1 {
			list.tail = currNode
		}
	}
	list.length--
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
