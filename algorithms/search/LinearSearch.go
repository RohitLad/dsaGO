package search

func LinearSearch(list []int, elem int) (int, int) {
	for i, v := range list {
		if v == elem {
			return v, i
		}
	}
	return 0, -1
}
