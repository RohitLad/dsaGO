package main

import (
	ds "dsaGO/datastructures"
)

func main() {
	testLinkedList()
	//testDoubleLinkedList()
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
