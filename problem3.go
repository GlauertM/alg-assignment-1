package main

// import "fmt"
// O(n)
func rotateArray(nums []int, left, right int) []int {
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
	return nums
}

// O(3 n) = O(n)
func partialRotation_of_anArray(nums []int, k int) []int {
	n := len(nums)

	rotateArray(nums, 0, n-1)
	rotateArray(nums, 0, k%n-1)
	rotateArray(nums, k%n, n-1)
	return nums
}

// func main() {
// 	sorted_array := []int{1, 2, 3, 6, 7, 8, 10, 11, 12, 16, 22}
// 	fmt.Println(partialRotation_of_anArray(sorted_array, 4))
// }
