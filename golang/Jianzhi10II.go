package main

import "math"

func numWays(n int) int {
	first, second := 1, 1
	if n == 0 {
		return first
	} else if n == 1 {
		return second
	}

	for i := 0; i < n-1; i++ {
		first, second = second, (first+second)%int(math.Pow10(9)+7)
	}
	return second
}
