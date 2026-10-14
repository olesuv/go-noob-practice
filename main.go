package main

import "fmt"

func main() {
	fmt.Println(max([]int{10, 5, 30, 20})) // 30 true
	fmt.Println(max([]int{-5, -2, -10}))   // -2 true
	fmt.Println(max([]int{}))              // 0 false
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
