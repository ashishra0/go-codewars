package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(PrinterError("kkkwwwaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbmmmmmmmmmmmmmmmmmmmxyz"))
}

func PrinterError(str string) string {
	total := len(str)
	var errors int
	words := strings.Split(str, "")
	for _, word := range words {
		if !ValidChars(word) {
			errors += 1
		}
	}
	return fmt.Sprintf("%d/%d", errors, total)
}

func ValidChars(str string) bool {
	return strings.ContainsAny(str, "abcdefghijklm")
}