package sort

import (
	"fmt"
)

func BubbleSort(list []int) {
	nElems := len(list)
	ifSwapped := true
	for ifSwapped {
		ifSwapped = false
		for i := 0; i < nElems-1; i++ {
			// bubble would be (list[i], list[i+1])
			if list[i+1] < list[i] {
				list[i], list[i+1] = list[i+1], list[i]
				ifSwapped = true
			}
		}
	}
}

func TestBubbleSort() {
	slice := []int{5, 7, 8, 10, 4, 2, 3}
	BubbleSort(slice)
	fmt.Println(slice)
}
