package main

import "fmt"

func main() {
	nums := []int{10, 20, 30}
	res := sum(nums)
	fmt.Println(res)
}

func sum(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}
