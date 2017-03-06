package main

import "fmt"

func isPalindrome(x int) bool {
	reverse_int := 0
	temp_int := x
	for temp_int > 0 {
		reverse_int = reverse_int*10 + temp_int%10
		temp_int /= 10
	}
	if (reverse_int == x) {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(isPalindrome(-1))
}