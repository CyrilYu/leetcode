package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	result := 0
	neg := true
	if (x < 0) {
		neg = false
		x *= -1
	}
	for x >= 10 {
		result = result * 10 + x%10 * 10
		x /= 10
	}
	result += x
	Max := math.MaxInt32
	Min := math.MinInt32
	if (result > Max || result < Min) {
		return 0
	}
	if (!neg) {
		result *= -1
	}
	return result
}

func main() {
	fmt.Println(reverse(-3984314))
}