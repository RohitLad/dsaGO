package main

import (
	ds "DSA/datastructures"
)

func main() {
	testLinkedList()
}

func testLinkedList() {
	list := ds.LinkedList{}
	list.Insert(55.0)
	list.Insert(88)
	list.Insert(68.125)
	list.Display()
}
