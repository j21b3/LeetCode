package main

import "math"

func printNumbers(n int) []int {
	limit := int(math.Pow10(n))
	ret := []int{}
	for i := 1; i < limit; i++ {
		ret = append(ret, i)
	}
	return ret
}
