package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(GetCount("abracadabra"))
}

func GetCount(str string) int {
	words := strings.Split(str, "")
	count := 0
	for _, word := range words {
		if IsVowel(word) {
			count += 1
		}
	}
	return count
}

func IsVowel(str string) bool {
	return strings.ContainsAny(str, "aeiou")
}