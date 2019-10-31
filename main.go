package main

import (
	ds "dsaGO/datastructures"
	"fmt"
)

func main() {
	testBinaryTree()
	//testQueue()
	//testStack()
	//testHashTable()
	//testLinkedList()
	//testDoubleLinkedList()
}

func testBinaryTree() {
	bt := ds.BinaryTree{}
	bt.Insert(50)
	bt.Insert(66)
	bt.Insert(22)
	bt.Insert(55)
	bt.Insert(1)
	bt.Insert(23)
	bt.Insert(68.2)
	bt.Display()
}

func testQueue() {
	queue := ds.Queue{}
	queue.Enqueue(55)
	queue.Enqueue(65.2)
	fmt.Println("First element: ", queue.Peek())
	queue.Enqueue("test")
	queue.Display()
	fmt.Println("Removed element: ", queue.Dequeue())
	queue.Display()
}

func testStack() {
	stack := ds.Stack{}
	stack.Push(55)
	stack.Push(65.2)
	fmt.Println("Top element: ", stack.Peek())
	stack.Push("test")
	stack.Display()
	fmt.Println("Popped element: ", stack.Pop())
	stack.Display()
}

func testHashTable() {
	table := ds.HashTable{}
	table.Set("Rohit", 25)
	table.Set("Rohito", 250)
	table.Set("A", 26.36)
	table.Set("H", 2645)
	table.Set("G", "chc")
	table.Set("E", 25.05)
	table.Set("Rohitoo", "test")
	table.Display()
	key2get := "Rohit"
	fmt.Println("Data at key: ", key2get, " is", table.Get(key2get))
	table.Remove(key2get)
	table.Display()
	fmt.Println("Keys: ", table.Keys())
}

func testDoubleLinkedList() {
	list := ds.DoubleLinkedList{}
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

func testLinkedList() {
	list := ds.LinkedList{}
	list.Append(55.0)
	list.Append(88)
	list.Append(68.125)
	list.Pretend(458)
	list.Display()
	list.Insert(2, 33.5)
	list.Display()
	list.Remove(4)
	list.Append(78)
	list.Remove(1)
	list.Display()
}
