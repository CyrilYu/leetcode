// Given an array S of n integers, find three integers in S 
// such that the sum is closest to a given number, 
// target. Return the sum of the three integers. You may assume that each input would have exactly one solution.

// For example, given array S = {-1 2 1 -4}, and target = 1.

// The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).

package main

import (
	"fmt"
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	result := nums[0] + nums[1] + nums[2]
	if (result == target) {
		return result
	}
	for i := 0; i < len(nums)-2; i++ {
		j := i+1
		k := len(nums)-1
		var sum int
		for j < k {
			sum = nums[i] + nums[j] + nums[k]
			if (sum == target) {
				result = sum
				break
			}
			if (sum < target) {
				j++
			} else if (sum > target) {
				k--
			}
			if (math.Abs(float64(sum - target)) < math.Abs(float64(result - target))) {
				result = sum
			}
		}
	}
	return result
}

func main() {
	arr := []int{-3,-2,-5,3,-4}
	fmt.Println(threeSumClosest(arr, -1))
}
