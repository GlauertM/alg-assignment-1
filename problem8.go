package main

func zeroLast(arr []int) []int {
	pntr_l := 0
	pntr_r := len(arr) - 1
	if len(arr) == 1 {
		return arr
	}

	for pntr_l <= pntr_r {
		if arr[pntr_r] == 0 {
			pntr_r--
		}
		if arr[pntr_l] == 0 {
			arr[pntr_l], arr[pntr_r] = arr[pntr_r], arr[pntr_l]
			pntr_r--

		}
		pntr_l++
	}
	return arr
}

// func main() {
// 	sortedArray := []int{0, 0}
// 	fmt.Println(zeroLast(sortedArray))
// }
