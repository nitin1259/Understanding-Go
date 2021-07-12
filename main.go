package main

import (
	"fmt"

	"github.com/nitin1259/Understanding-Go/dynpro"
	// "github.com/nitin1259/Understanding-Go/dynpro"
)

func main() {

	weights := [4]int{1, 3, 4, 5}
	values := [4]int{1, 4, 5, 7}

	capacity := 7

	// create 2d slice
	t := make([][]int, len(weights))
	for i := range t {
		t[i] = make([]int, capacity)
	}

	// initializing the memoization matrix with -1
	for i, tt := range t {
		for j, _ := range tt {
			t[i][j] = -1
		}
	}

	// fmt.Println(weights)
	// fmt.Println(values)
	// fmt.Println(capacity)
	fmt.Println(dynpro.FindMax(weights, values, capacity, 4, t))

	for i, tt := range t {
		fmt.Println()
		for j, _ := range tt {
			fmt.Print(t[i][j], " ")
		}
	}
}
