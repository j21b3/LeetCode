package main

func cuttingRope(n int) int {
	if n <= 3 {
		return n - 1
	}
	a, b := int(n/3), n%3
	if b == 0 {
		return quickPow(3, a)
	} else if b == 1 {
		return (quickPow(3, a-1) * 4) % 1000000007
	} else {
		return (quickPow(3, a) * 2) % 1000000007
	}
}

func quickPow(a, n int) int { // a^n
	if n == 0 {
		return 1
	} else if n == 1 {
		return a
	}

	if n%2 == 0 {
		half := quickPow(a, n/2)
		return (half * half) % 1000000007
	} else {
		half := quickPow(a, (n-1)/2)
		return (half * half * a) % 1000000007
	}
}
