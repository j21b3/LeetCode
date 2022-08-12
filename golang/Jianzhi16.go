package main

// 递归写法
/* func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	} else if n < 0 {
		return 1 / myPow(x, -n)
	}

	if n%2 == 0 {
		ret := myPow(x, n/2)
		return ret * ret
	} else {
		ret := myPow(x, (n-1)/2)
		return ret * ret * x
	}
} */

// 非递归
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	} else if n < 0 {
		x, n = 1/x, -n
	}
	ret := 1.0
	for n != 0 {
		if n&1 == 1 {
			ret = ret * x
		}
		x = x * x
		n = n >> 1
	}
	return ret
}
