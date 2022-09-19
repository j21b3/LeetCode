package main

func majorityElement(nums []int) int {
	selectnum, count := 0, 0
	for _, each := range nums {
		if count == 0 {
			selectnum = each
			count = 1
		} else if selectnum == each {
			count++
		} else {
			count--
		}
	}
	return selectnum
}
