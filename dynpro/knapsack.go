package dynpro

import "math"

func FindMax(weight, value [4]int, capacity int, n int, t [][]int) int {

	// base condiditon
	if n == 0 || capacity == 0 {
		return 0
	}

	// checking with memoization, if it resolved earlier
	if t[n-1][capacity-1] != -1 {
		return t[n-1][capacity-1]
	}

	if weight[n-1] <= capacity {
		// return int(math.Max(float64(value[n-1]+FindMax(weight, value, capacity-weight[n-1], n-1, t)), float64(FindMax(weight, value, capacity, n-1, t))))
		t[n-1][capacity-1] = int(math.Max(float64(value[n-1]+FindMax(weight, value, capacity-weight[n-1], n-1, t)), float64(FindMax(weight, value, capacity, n-1, t))))
		return t[n-1][capacity-1]
	} else {
		// return FindMax(weight, value, capacity, n-1, t)
		return t[n-1][capacity-1]
	}

}
