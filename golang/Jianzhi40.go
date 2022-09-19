package main

func getLeastNumbers(arr []int, k int) []int {
	// heap里存堆的先序遍历
	maxheap := make([]int, k)
	copy(maxheap, arr[:k])

	var adjust func(idx int)
	adjust = func(idx int) {
		biggest := idx
		//如果idx比左孩子小
		if 2*idx+1 < k && maxheap[biggest] < maxheap[2*idx+1] {
			biggest = 2*idx + 1
		}
		//如果idx比右孩子小
		if 2*idx+2 < k && maxheap[biggest] < maxheap[2*idx+2] {
			biggest = 2*idx + 2
		}
		if biggest != idx {
			maxheap[biggest], maxheap[idx] = maxheap[idx], maxheap[biggest]
			adjust(biggest)
		}
	}

	gen := func() {
		for i := int(k / 2); i > -1; i-- {
			adjust(i)
		}
	}

	if k == 0 {
		return []int{}
	}
	gen()
	for i := k; i < len(arr); i++ {
		if arr[i] < maxheap[0] {
			maxheap[0] = arr[i]
			gen()
		}
	}
	return maxheap
}
