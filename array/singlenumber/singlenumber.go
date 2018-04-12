package main

func SingleNumber(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	result := 0
	for _, n := range nums {
		result ^= n
	}
	return result
}
