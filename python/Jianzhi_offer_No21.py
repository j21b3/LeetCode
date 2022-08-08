class Solution:
    def exchange(self, nums):
        i, j = 0, len(nums) - 1
        while i <= j and i < len(nums) and j >= 0:
            if nums[i] % 2 == 0 and nums[j] % 2 == 1:
                nums[i], nums[j] = nums[j], nums[i]
            if nums[i] % 2 == 1:
                i += 1
            if nums[j] % 2 == 0:
                j -= 1
        return nums


s = Solution()
print(s.exchange([1, 2, 3, 4]))
