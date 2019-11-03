package sort

import (
	"fmt"
	"math/rand"
	orgsrt "sort"
)

func QuickSort(list []int) {

	if len(list) <= 1 {
		return
	}
	p := pivot(list)
	left, right := 0, len(list)-1
	list[left], list[p] = list[p], list[left]
	for left != right {
		if list[left] < list[left+1] {
			list[left+1], list[right] = list[right], list[left+1]
			right--
		} else {
			list[left+1], list[left] = list[left], list[left+1]
			left++
		}
	}
	QuickSort(list[:left])
	QuickSort(list[left+1:])
}

func pivot(list []int) int {
	//return rand.Int() % len(list)
	return len(list) - 1
}

func TestQuickSort() {
	N := 200
	arr1 := make([]int, N)
	arr2 := make(orgsrt.IntSlice, N)
	randomSlice(arr1)
	copy(arr2, arr1)
	QuickSort(arr1)
	arr2.Sort()
	testSameness(arr1, arr2)
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
