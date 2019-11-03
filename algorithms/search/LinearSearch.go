package search

import "fmt"

func LinearSearch(list []int, elem int) (int, int) {
	for i, v := range list {
		if v == elem {
			return v, i
		}
	}
	return 0, -1
}

func TestLinearSearch() {
	num2Find := 8
	slice := []int{5, 7, 8, 10, 4, 2, 3}
	_, i := LinearSearch(slice, num2Find)
	fmt.Println(num2Find, " is present at index ", i, " of slice")
}
