package sort

import "fmt"

func InsertionSort(list []int) {

	for i, v := range list {
		for i > 0 {
			if list[i-1] > v {
				list[i] = list[i-1]
				i--
			} else {
				break
			}
		}
		list[i] = v
	}
}

func TestInsertionSort() {
	slice := []int{5, 7, 8, 10, 4, 2, 3}
	InsertionSort(slice)
	fmt.Println(slice)
}
