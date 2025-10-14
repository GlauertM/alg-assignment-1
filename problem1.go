package main

//import "fmt"

func two_sum(nums []int, target int) []int {
	left := 0
	right := len(nums) - 1

	for left < right {
		sum := nums[left] + nums[right]

		if sum == target {
			return []int{nums[left], nums[right]}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return []int{nums[left], nums[right]}
}

// func main() {
// 	sorted_array := []int{1, 2, 3, 6, 7, 8, 10, 11, 12, 16, 22}
// 	fmt.Println(two_sum(sorted_array, 13))
// }
