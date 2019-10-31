package sort

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
