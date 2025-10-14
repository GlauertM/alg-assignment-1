package main

func zeroLast(arr []int) []int {
	pntr_l := 0
	pntr_r := len(arr) - 1
	// for pntr_l <= pntr_r {
	// 	if arr[pntr_l] == 0 {
	// 		arr[pntr_l], arr[pntr_r] = arr[pntr_r], arr[pntr_l]
	// 		pntr_r--
	// 		if arr[pntr_l] == 0 {
	// 			pntr_l--
	// 		}
	// 	}
	// 	pntr_l++
	// }

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
// 	sorted_array := []int{0, 0, 1, 2, 0, 6, 0, 8, 10, 0, 12, 16, 22}
// 	fmt.Println(zeroLast(sorted_array))
// }
