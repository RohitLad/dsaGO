package datastructures

import (
	"fmt"
	"hash/fnv"
)

const hashtablesize int = 5

type hnode struct {
	Data interface{}
	key  string
	next *hnode
}

type HashTable struct {
	data [hashtablesize]*hnode
}

func (ht HashTable) generateHash(key string) uint8 {
	h := fnv.New32a()
	h.Write([]byte(key))
	return uint8(h.Sum32() % uint32(hashtablesize))
}

func (ht *HashTable) Set(key string, elem interface{}) {
	ind := ht.generateHash(key)
	if ht.data[ind] == nil {
		newNode := hnode{Data: elem, key: key}
		ht.data[ind] = &newNode
	} else {
		currNode := ht.data[ind]
		for currNode.next != nil {
			if currNode.key == key {
				currNode.Data = elem
				return
			} else {
				currNode = currNode.next
			}
		}
		newNode := hnode{Data: elem, key: key}
		currNode.next = &newNode
	}

}

func (ht HashTable) Get(key string) interface{} {

	currNode, _, _ := ht.goTo(key)

	if currNode != nil {
		return currNode.Data
	} else {
		return nil
	}
}

func (ht *HashTable) Remove(key string) {
	currNode, prevNode, ind := ht.goTo(key)
	if ht.data[ind] == currNode {
		ht.data[ind] = currNode.next
	} else if currNode != nil {
		nextNode := currNode.next
		prevNode.next = nextNode
	}
}

func (ht *HashTable) goTo(key string) (*hnode, *hnode, uint8) {
	ind := ht.generateHash(key)
	currNode := ht.data[ind]
	prevNode := &hnode{}
	for currNode != nil {
		if currNode.key == key {
			return currNode, prevNode, ind
		} else {
			prevNode = currNode
			currNode = currNode.next
		}
	}
	return currNode, nil, uint8(0)
}

func (ht HashTable) Display() {
	for i := 0; i < hashtablesize; i++ {
		fmt.Printf("|")
		currNode := ht.data[i]
		for currNode != nil {
			fmt.Printf("--->  %v  ", currNode.Data)
			currNode = currNode.next
		}
		fmt.Printf("\n")
	}
}
