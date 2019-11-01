package main

import (
	srch "dsaGO/algorithms/search"
	sort "dsaGO/algorithms/sort"
	ds "dsaGO/datastructures"
	"fmt"
	"math/rand"
	orgsrt "sort"
)

func main() {
	testQuickSort()
	//testMergeSort()
	//testInsertionSort()
	//testBubbleSort()
	//testLinearSearch()
	//testBinaryTree()
	//testQueue()
	//testStack()
	//testHashTable()
	//testLinkedList()
	//testDoubleLinkedList()
}

func testSameness(list1 []int, list2 []int) {
	if len(list1) != len(list2) {
		fmt.Println("Trying to compare lists of unequal size")
		return
	}
	for i := range list1 {
		if list1[i] != list2[i] {
			fmt.Println("problem at index ", i, " of list 1")
			break
		}
	}
	fmt.Println("Both lists are same")
}

func randomSlice(list []int) {
	N := len(list)
	for i := range list {
		list[i] = rand.Intn(N)
	}
}

func testQuickSort() {
	N := 200
	arr1 := make([]int, N)
	arr2 := make(orgsrt.IntSlice, N)
	randomSlice(arr1)
	copy(arr2, arr1)
	sort.QuickSort(arr1)
	arr2.Sort()
	testSameness(arr1, arr2)
}

func testMergeSort() {
	slice := []int{5, 7, 8, 10, 4, 2, 3}
	sort.MergeSort(slice)
	fmt.Println(slice)
}

func testInsertionSort() {
	slice := []int{5, 7, 8, 10, 4, 2, 3}
	sort.InsertionSort(slice)
	fmt.Println(slice)
}

func testBubbleSort() {
	slice := []int{5, 7, 8, 10, 4, 2, 3}
	sort.BubbleSort(slice)
	fmt.Println(slice)
}

func testLinearSearch() {
	num2Find := 8
	slice := []int{5, 7, 8, 10, 4, 2, 3}
	_, i := srch.LinearSearch(slice, num2Find)
	fmt.Println(num2Find, " is present at index ", i, " of slice")
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

func testDoubleLinkedList() ds.DoubleLinkedList {
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
	return list
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
