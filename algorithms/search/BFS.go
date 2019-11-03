package search

import (
	ds "dsaGO/datastructures"
	"fmt"
)

func BFS(tree ds.BinaryTree, elem interface{}) bool {
	if tree.Root == nil {
		return false
	}
	root := tree.Root()
	searchQueue := ds.Queue{}
	searchQueue.Enqueue(root)
	for searchQueue.Size() != 0 {
		currNode := searchQueue.Dequeue().(*ds.BTNode)
		if currNode.Value() == elem {
			return true
		}
		if currNode.Left() != nil {
			searchQueue.Enqueue(currNode.Left())
		}
		if currNode.Right() != nil {
			searchQueue.Enqueue(currNode.Right())
		}
	}
	return false
}

func TestBFS() {
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
	ok := BFS(bt, elem)
	if ok {
		fmt.Println(elem, " exists in the BT")
	} else {
		fmt.Println(elem, " does not exist in the BT")
	}

}
