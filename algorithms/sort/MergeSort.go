package sort

import "fmt"

func mDivide(array []int, lIndex int, rIndex int) []int {
	mid := (lIndex + rIndex) / 2
	diff := rIndex - lIndex
	if diff/2 == 0 {
		return []int{array[lIndex]}
	}
	lArr := mDivide(array, lIndex, mid)
	rArr := mDivide(array, mid, rIndex)
	tmpArr := []int{}
	lCount := 0
	rCount := 0
	for i := 0; i < len(lArr)+len(rArr); i++ {
		if lCount == len(lArr) {
			tmpArr = append(tmpArr, rArr[rCount:]...)
			break
		} else if rCount == len(rArr) {
			tmpArr = append(tmpArr, lArr[lCount:]...)
			break
		} else {
			if lArr[lCount] > rArr[rCount] {
				tmpArr = append(tmpArr, rArr[rCount])
				rCount++
			} else {
				tmpArr = append(tmpArr, lArr[lCount])
				lCount++
			}
		}
	}
	return tmpArr
}

func MergeSort(list []int) {
	sorted := mDivide(list, 0, len(list))
	for i, _ := range list {
		list[i] = sorted[i]
	}
}

func TestMergeSort() {
	slice := []int{5, 7, 8, 10, 4, 2, 3}
	MergeSort(slice)
	fmt.Println(slice)
}
