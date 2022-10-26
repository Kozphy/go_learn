package week1

import "fmt"

func search(nums []int, target int) int {
	low := 0
	high := len(nums) - 1
	for low <= high {
		mid := (low + high) / 2

		if nums[mid] == target {
			return mid
		}

		if target < (nums[low]+nums[high])/2 {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

func Execute_search() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 9
	ans := search(nums, target)
	fmt.Println(ans)
}
