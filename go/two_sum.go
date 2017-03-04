package main

import "fmt"

func sum(nums []int, target int) []int {
	result := []int{}
	index := 0
	for i := 0; i < len(nums); i++ {
		for index = i + 1; index < len(nums); index++ {
			if nums[i]+nums[index] == target {
				result = []int{i, index}
			} else {
				continue
			}
		}

	}
	return result
}

func main() {
	nums := []int{3, 2, 3}
	target := 6
	fmt.Println(sum(nums, target))
}
