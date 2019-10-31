package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(FindShort("Welcome to the world of go"))
}

func FindShort(s string) int {
	words := strings.Split(s, " ")
	shortest := words[0]

	for _, word := range words {
		if len(word) < len(shortest) {
			shortest = word
		}
	}
	return len(shortest)
}