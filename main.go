package main

import "fmt"

func main() {
	fmt.Println(findIndex([]int{10, 20, 30}, 20)) // 1
	fmt.Println(findIndex([]int{10, 20, 30}, 50)) // -1
	fmt.Println(findIndex([]int{}, 10))           // -1
}

func findIndex(nums []int, target int) int {
	for index, num := range nums {
		if num == target {
			return index
		}
	}
	return -1
}
