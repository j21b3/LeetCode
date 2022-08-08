package main

/*
* 每次将nums[i]的元素对应其值所对应的位置，当放入已经放过的地方就说明重复了
 */

func findRepeatNumber(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for nums[i] != i {
			if nums[nums[i]] == nums[i] {
				return nums[i]
			}
			nums[nums[i]], nums[i] = nums[i], nums[nums[i]]
		}
	}
	return -1
}
