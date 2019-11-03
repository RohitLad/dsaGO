package search

import (
	ds "dsaGO/datastructures"
	"fmt"
)

func DFS(tree ds.BinaryTree, elem interface{}) bool {
	if tree.Root == nil {
		return false
	}

	list := ds.LinkedList{}
	bul := false
	callback := func(val interface{}) {
		list.Append(val)
		if val == elem {
			bul = true
		}
	}
	//TraverseInOrder(tree.Root(), callback)
	//TraversePreOrder(tree.Root(), callback)
	TraversePostOrder(tree.Root(), callback)
	fmt.Printf("Traversal: ")
	list.Display()
	return bul
}

func TraverseInOrder(node *ds.BTNode, callback func(interface{})) {
	if node.Left() != nil {
		TraverseInOrder(node.Left(), callback)
	}
	callback(node.Value())
	if node.Right() != nil {
		TraverseInOrder(node.Right(), callback)
	}
}

func TraversePreOrder(node *ds.BTNode, callback func(interface{})) {
	callback(node.Value())
	if node.Left() != nil {
		TraversePreOrder(node.Left(), callback)
	}
	if node.Right() != nil {
		TraversePreOrder(node.Right(), callback)
	}
}

func TraversePostOrder(node *ds.BTNode, callback func(interface{})) {
	if node.Left() != nil {
		TraversePostOrder(node.Left(), callback)
	}
	if node.Right() != nil {
		TraversePostOrder(node.Right(), callback)
	}
	callback(node.Value())
}

func TestDFS() {
	bt := ds.BinaryTree{}
	bt.Insert(50)
	bt.Insert(66)
	bt.Insert(22)
	bt.Insert(55)
	bt.Insert(1)
	elem := 23
	bt.Insert(elem)
	bt.Insert(68.2)
	bt.Display()
	ok := DFS(bt, elem)
	if ok {
		fmt.Println(elem, " exists in the BT")
	} else {
		fmt.Println(elem, " does not exist in the BT")
	}

}
