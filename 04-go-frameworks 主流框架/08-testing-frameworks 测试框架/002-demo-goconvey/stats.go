package main

func Average(nums []int) float64 {
	if len(nums) == 0 {
		return 0
	}

	total := 0
	for _, num := range nums {
		total += num
	}
	return float64(total) / float64(len(nums))
}

func Max(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	max := nums[0]
	for _, num := range nums[1:] {
		if num > max {
			max = num
		}
	}
	return max
}
