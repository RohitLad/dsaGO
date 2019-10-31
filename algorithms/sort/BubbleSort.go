package sort

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
