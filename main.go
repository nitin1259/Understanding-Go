package main

import (
	"fmt"

	"github.com/nitin1259/Understanding-Go/dynpro"
)

func init() {

}

func main() {

	weights := [4]int{1, 3, 4, 5}
	values := [4]int{1, 4, 5, 7}

	capacity := 7

	fmt.Println(dynpro.FindMax(weights, values, capacity, 4))
}
