package main

import "fmt"

func twoSum(nums []int, target int) ([]int, error) {
	if nums == nil {
		return nil, fmt.Errorf("The array is empty.")
	}

	left := 0
	right := len(nums) - 1

	for left < right {
		sum := nums[left] + nums[right]

		if sum == target {
			return []int{left, right}, fmt.Errorf("")
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return nil, fmt.Errorf("Appropriate pair is not found.")
}

func main() {
	sortedArray := []int{1, 2, 3, 6, 7, 8, 10, 11, 12, 16, 22}
	nilArray := []int{}
	fmt.Println(twoSum(sortedArray, 13))
	fmt.Println(twoSum(nilArray, 16))
}
