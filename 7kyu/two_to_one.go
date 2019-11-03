package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	s1 := "loopingisfunbutdangerous"
	s2 := "lessdangerousthancoding"
	fmt.Println(TwoToOne(s1, s2))
}

func TwoToOne(s1 string, s2 string) string {
	var s3 string
	var result string
	s3 = s1 + s2
	sliceOfS3 := strings.Split(s3, "")
	sort.Strings(sliceOfS3)
	for _, str := range sliceOfS3 {
		if !strings.Contains(result, str) {
			result += str
		}
	}
	return result
}
