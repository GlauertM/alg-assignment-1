package main

func sortBinaryArray(arr []byte) []byte {
	pointer_left := 0
	pointer_right := len(arr) - 1
	for pointer_left < pointer_right {
		if arr[pointer_left] == 0 {
			pointer_left++
		} else if arr[pointer_right] == 1 {
			pointer_right--
		} else {
			arr[pointer_left], arr[pointer_right] = arr[pointer_right], arr[pointer_left]
			pointer_left++
			pointer_right--
		}
	}
	return arr
}

// func main() {
// 	binary_array := []byte{1, 0, 0, 1, 0, 1, 0, 1, 0, 0, 1, 1}
// 	fmt.Println(sortBinaryArray(binary_array))
// }
