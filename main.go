package main

import "fmt"

func main() {
	fmt.Println(sumEven([]int{1, 2, 3, 4, 5, 6})) // 12
	fmt.Println(sumEven([]int{1, 3, 5}))          // 0
	fmt.Println(sumEven([]int{}))                 // 0
}

func sumEven(nums []int) int {
	sum := 0

	for _, num := range nums {
		if num%2 == 0 {
			sum += num
		}
	}

	return sum
}
