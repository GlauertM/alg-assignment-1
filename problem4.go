package main

func mergeSortedArrays(arr1, arr2 []int) []int {
	pntr1 := len(arr1) - len(arr2) - 1
	pntr2 := len(arr2) - 1
	pntrMid := len(arr1) - 1
	for pntr2 >= 0 {
		if pntr1 >= 0 && arr1[pntr1] < arr2[pntr2] {
			arr1[pntrMid] = arr1[pntr1]
			pntr1--
		} else {
			arr1[pntrMid] = arr2[pntr2]
			pntr2--
		}
		pntrMid--
	}
	return arr1
}

// func main() {
// 	sortedArray1 := []int{1, 2, 3, 6, 7, 8, 10, 11, 12, 16, 22, 0, 0, 0, 0}
// 	sortedArray2 := []int{4, 5, 9, 13}
// 	fmt.Println(mergeSortedArrays(sortedArray1, sortedArray2))
// }
