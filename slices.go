package main

import (
	"fmt"
	"slices"
)

func main() {
	test_1 := []int{1, 5, 7, 2, 1, 2}
	test_2 := []string{"aku", "suka", "kamu"}

	fmt.Println(slices.Max(test_1))
	fmt.Println(slices.Min(test_1))
	fmt.Println(slices.Contains(test_1, 4))
	fmt.Println(slices.Contains(test_1, 2))
	fmt.Println(slices.Max(test_2))
	fmt.Println(slices.Min(test_2))
	fmt.Println(slices.Contains(test_2, "suka"))

}
