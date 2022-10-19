package array_hashing

import "fmt"

func containsDuplicate(nums []int) bool {
	mapping := make(map[int]int)
	for _, num := range nums {
		if mapping[num] != 0 {
			return true
		} else {
			mapping[num] += 1
		}
	}
	return false
}

func Execute_containsDuplicate() {
	nums := []int{1, 2, 3, 1}
	ans := containsDuplicate(nums)
	fmt.Println(ans)
}
