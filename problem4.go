package main

import "fmt"

func merge_sorted_arrays(arr1, arr2 []int) []int {
	// merged_array := []int{}
	// i, j := 0, 0

	// for i < len(arr1) && j < len(arr2) {
	// 	if arr1[i] < arr2[j] {
	// 		merged_array.append(merged_array, arr1[i])
	// 		i++
	// 	} else {
	// 		merged_array.append(merged_array, arr2[j])
	// 		j++
	// 	}
	// }

	// merged_array.append(merged_array, arr1[i])
	// merged_arrayappend(merged_array, arr2[j])

	// return merged_array

	pntr_arr2 := len(arr2) - 1
	last_pntr_arr1 := len(arr1) - 1

	for pntr_arr1 := len(arr1) - 1; pntr_arr2 >= 0; last_pntr_arr1-- {
		if pntr_arr1 >= 0 && arr2[pntr_arr2] < arr1[pntr_arr1] {
			arr1[last_pntr_arr1] = arr1[pntr_arr1]
			pntr_arr1--
		} else {
			arr1[last_pntr_arr1] = arr2[pntr_arr2]
			pntr_arr2--
		}
	}
	return arr1
}

func main() {
	sorted_array1 := []int{1, 2, 3, 6, 7, 8, 10, 11, 12, 16, 22, 0, 0, 0, 0}
	sorted_array2 := []int{4, 5, 9, 13}
	fmt.Println(merge_sorted_arrays(sorted_array1, sorted_array2))
}
