package main

func evenFirst(arr []int) []int {
	evenIndex := 0
	for i := 0; i < len(arr); i++ {
		if arr[i]%2 == 0 {
			arr[i], arr[evenIndex] = arr[evenIndex], arr[i]
			evenIndex++
		}
	}
	return arr
}

// func main() {
// 	sortedArray := []int{1, 2, 3, 6, 7, 8, 10, 11, 12, 16, 22}
// 	fmt.Println(evenFirst(sortedArray))
// }
