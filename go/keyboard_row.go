package main

import (
	"fmt"
	"strings"
)

func findWords(words []string) []string {
	var result []string
	letters := []string{"asdfghjkl", "qwertyuiop", "zxcvbnm"} 

	for i := 0; i < len(words); i++ {
		for num := range letters {
			str := letters[num]
			count := 0
			for index := range words[i] {
				temp_str := strings.ToLower(string(words[i][index]))
				if (strings.Index(str, temp_str) == -1) {
					break
				} else {
					count++
					if (count == len(words[i])) {
						result = append(result, words[i])
						break
					}
				}
			}
		}
	}
    return result
}

func main() {
	arr := []string{"Hello","Alaska","Dad","Peace"}
	fmt.Println(findWords(arr))
}