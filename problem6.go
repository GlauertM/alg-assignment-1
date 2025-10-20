package main

func sortColours(arr []byte) []byte {
	pntr_low := 0
	pntr_mid := 0
	pntr_high := len(arr) - 1

	for pntr_mid <= pntr_high {
		if arr[pntr_mid] == 0 {
			arr[pntr_low], arr[pntr_mid] = arr[pntr_mid], arr[pntr_low]
			pntr_low++
			pntr_mid++
		} else if arr[pntr_mid] == 1 {
			pntr_mid++
		} else if arr[pntr_mid] == 2 {
			arr[pntr_mid], arr[pntr_high] = arr[pntr_high], arr[pntr_mid]
			pntr_high--
		}
	}
	return arr
}

// func main() {
// 	binaryArray := []byte{2, 2, 1, 0, 2, 0, 1, 0, 1, 2, 1, 2, 0, 1, 1}
// 	fmt.Println(sortColours(binaryArray))
// }
