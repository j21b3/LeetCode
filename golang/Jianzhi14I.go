package main

import "math"

/*
*	思路见 https://leetcode.cn/problems/jian-sheng-zi-lcof/solution/mian-shi-ti-14-i-jian-sheng-zi-tan-xin-si-xiang-by/
 */

func cuttingRope(n int) int {
	if n <= 3 {
		return n - 1
	}

	a, b := int(n/3), n%3
	if b == 0 {
		return int(math.Pow(3, float64(a)))
	} else if b == 1 {
		return int(math.Pow(3, float64(a-1))) * 4
	} else {
		return int(math.Pow(3, float64(a))) * 2
	}
}
