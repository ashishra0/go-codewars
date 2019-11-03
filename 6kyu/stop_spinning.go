package main

import (
	"fmt"
	"strings"
)


func main() {
	fmt.Println(SpinWords("Hey spinning words"))
}

func SpinWords(str string) string {
	var result []string
	newStr := strings.Split(str, " ")
	for _, word := range newStr {
		if len(word) >= 5 {
			word = reverse(word)
		}
		result = append(result, word)
	}
	return strings.Join(result, " ")
}


func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}