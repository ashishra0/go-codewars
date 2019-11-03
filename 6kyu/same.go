package main

import (
	"fmt"
	"reflect"
	"sort"
)

func main() {
	a1 := []int{121, 144, 19, 161, 19, 144, 19, 11}
	a2 := []int{11*21, 121*121, 144*144, 19*19, 161*161, 19*19, 144*144, 19*19}
	fmt.Println(Comp(a1, a2))
}

func Comp(array1 []int, array2 []int) bool {
	var newArr1 []int
	var newArr2 []int

	for _, num := range array1 {
		newNum := num * num
		newArr1 = append(newArr1, newNum)
	}

	for _, num := range array2 {
		newArr2 = append(newArr2, num)
	}

	sort.Ints(newArr1)
	sort.Ints(newArr2)
	if reflect.DeepEqual(newArr1, newArr2) {
		return true
	}
	return false
}