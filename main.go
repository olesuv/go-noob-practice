package main

import "fmt"

func main() {
	num, bool_val := max([]int{10, 20, 30})
	fmt.Println(num, bool_val)
}

func max(nums []int) (int, bool) {
	if len(nums) == 0 {
		return 0, false
	}

	max := nums[0]

	for _, num := range nums {
		if num > max {
			max = num
		}
	}

	return max, true
}
