package datastructures

import (
	"errors"
	"fmt"
)

type btnode struct {
	left  *btnode
	right *btnode
	value interface{}
}

func (btn btnode) comparable() (float64, bool) {
	flt, ok := btn.value.(float64)
	if ok {
		return flt, ok
	}
	in, ok := btn.value.(int)
	if ok {
		return float64(in), ok
	}
	return 0., false
}

type BinaryTree struct {
	root   *btnode
	levels int
}

func (bt *BinaryTree) Insert(elem interface{}) {
	newNode := btnode{value: elem}

	if bt.root == nil {
		bt.root = &newNode
		bt.levels++
	} else {
		currNode := bt.root
		count := 1
		breakFlag := false
		for !breakFlag {
			el1, ok1 := newNode.comparable()
			el2, ok2 := currNode.comparable()
			if ok1 && ok2 {
				if el1 <= el2 {
					if currNode.left == nil {
						currNode.left = &newNode
						breakFlag = true
					}
					currNode = currNode.left
				} else {
					if currNode.right == nil {
						currNode.right = &newNode
						breakFlag = true
					}
					currNode = currNode.right
				}
				count++
				if breakFlag {
					maxlevelsexceeded := count > bt.levels
					if maxlevelsexceeded {
						bt.levels = count
					}
				}
			} else {
				err := errors.New("Cannot compare elements")
				fmt.Println(err)
				break
			}
		}

	}
}

func (bt BinaryTree) Display() {

}
