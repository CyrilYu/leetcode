package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
    m := make(map[int]int)
    for i, n := range nums {
		fmt.Println(m[n])
		fmt.Println(i)
		_, prs := m[n]
		fmt.Println(prs)
        if prs {
            return []int{m[n], i}
        } else {
            m[target-n] = i
        }
    }
    return nil
}

func main() {
	num := []int{-10,7,19,15, 22}
	target := 22
	twoSum(num, target)
	fmt.Println("Test")
}