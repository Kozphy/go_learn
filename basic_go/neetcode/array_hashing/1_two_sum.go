package array_hashing

import "fmt"

func twoSum(nums []int, target int) []int {
	mapping_nums := make(map[int]int)

	for pos, num := range nums {
		mapping_nums[num] = pos
	}

	for i := 0; i < len(nums); i++ {
		want := target - nums[i]
		if mapping_nums[want] != 0 && i != mapping_nums[want] {
			return []int{i, mapping_nums[want]}
		}
	}
	return []int{}

}

func twoSum_2(nums []int, target int) []int {
	return []int{}
}

func Execute_twoSum() {
	nums := []int{2, 7, 11, 15}
	target := 9
	ans := twoSum(nums, target)
	fmt.Println(ans)

}
