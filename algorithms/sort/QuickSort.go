package sort

func QuickSort(list []int) {

	if len(list) <= 1 {
		return
	}
	p := partition(list)
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

func partition(list []int) int {
	return len(list) - 1 //rand.Int() % len(list)
}