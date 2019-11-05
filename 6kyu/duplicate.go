package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(duplicateCount("abcdea1B1"))
}

func duplicateCount(str string) int {
	var duplicate []string
	str = strings.ToLower(str)
	words := strings.Split(str, "")
	for _, word := range words {
		if strings.Count(str, word) > 1 {
			duplicate = append(duplicate, word)
		}
	}
	duplicate = uniqueSlice(duplicate)
	return len(duplicate)
}

func uniqueSlice(arr []string) []string{
	unique := make(map[string]bool)
	var us []string
	for _, elem := range arr {
		if !unique[elem] {
			us = append(us, elem)
			unique[elem] = true
		}
	}
	return us
}