package main

import "fmt"

func main() {
	fmt.Println(countOccurrences([]int{1, 2, 2, 3, 2}, 2)) // 3
	fmt.Println(countOccurrences([]int{1, 2, 3}, 5))       // 0
	fmt.Println(countOccurrences([]int{}, 1))              // 0
}

func countOccurrences(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	counter := 0

	for _, num := range nums {
		if num == target {
			counter += 1
		}
	}

	return counter
}
