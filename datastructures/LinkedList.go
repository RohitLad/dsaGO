package datastructures

import "fmt"

type LinkedList struct {
	Head *Node
}

type Node struct {
	Data interface{}
	Next *Node
}

func (list *LinkedList) Insert(elem interface{}) {
	node := Node{Data: elem}
	if list.Head == nil {
		list.Head = &node
	} else {
		currNode := list.Head
		for currNode.Next != nil {
			currNode = currNode.Next
		}
		currNode.Next = &node
	}
}

func (list LinkedList) Display() {

	node := list.Head
	if node == nil {
		fmt.Println("Empty Linked List")
	} else {
		for node != nil {
			fmt.Printf("%v ->", node.Data)
			node = node.Next
		}
		fmt.Println("")
	}
}
