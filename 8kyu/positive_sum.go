package main

import "fmt"

func main() {
	fmt.Println(PositiveSum([]int{1, 2, 3, -4, 5}))
}

func PositiveSum(numbers []int) int {
	positive := []int{0}
	for _, number := range numbers {
		if number > 0 {
			positive = append(positive, number)
		}
	}
	sum := 0
	for _, num := range positive {
		sum += num
	}
	return sum
}
