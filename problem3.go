package main

import "fmt"

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
func partialRotationOfArray(nums []int, k int) ([]int, error) {
	n := len(nums)
	// Checking if the array 'nums' is empty.
	if n > 0 {
		rotateArray(nums, 0, n-1)
		rotateArray(nums, 0, k%n-1)
		rotateArray(nums, k%n, n-1)
		return nums, fmt.Errorf("")
	} else {
		return nil, fmt.Errorf("Array is empty.")
	}

}

func main() {
	sortedArray := []int{}
	fmt.Println(partialRotationOfArray(sortedArray, 4))
}
