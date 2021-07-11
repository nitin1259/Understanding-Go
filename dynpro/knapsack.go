package dynpro

import "math"

func FindMax(weight, value [4]int, capacity int, n int) int {

	// base condiditon
	if n == 0 || capacity == 0 {
		return 0
	}

	if weight[n-1] <= capacity {
		return int(math.Max(float64(value[n-1]+FindMax(weight, value, capacity-weight[n-1], n-1)), float64(FindMax(weight, value, capacity, n-1))))
	} else {
		return FindMax(weight, value, capacity, n-1)
	}

}
