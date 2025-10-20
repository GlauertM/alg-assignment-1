package main

// O(n) Linear time
func reverseArray(nums []int) []int {
	left := 0
	right := len(nums) - 1

	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
	return nums
}

// func main() {
// 	sortedArray := []int{1, 2, 3, 6, 7, 8, 10, 11, 12, 16, 22}
// 	fmt.Println(reverseArray(sortedArray))
// }
