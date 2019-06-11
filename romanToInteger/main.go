/**
I can be placed before V (5) and X (10) to make 4 and 9.
X can be placed before L (50) and C (100) to make 40 and 90.
C can be placed before D (500) and M (1000) to make 400 and 900.
**/

package main

import "fmt"

var num = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func romanToInt(s string) int {
	out := 0
	ln := len(s)
	for i := 0; i < ln; i++ {
		c := string(s[i])
		vc := num[c]
		if i < ln-1 {
			cnext := string(s[i+1])
			vcnext := num[cnext]
			if vc < vcnext {
				out += vcnext - vc
				i++
			} else {
				out += vc
			}
		} else {
			out += vc
		}
	}
	return out
}

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}
