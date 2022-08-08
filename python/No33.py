class Solution:
    def search(self, nums: List[int], target: int) -> int:
        if nums[0]<=target:
            i=0
            while i<len(nums) and nums[i]>=nums[0]:
                if nums[i]==target:
                    return i
                i+=1
        else:
            i=len(nums)-1
            while i>=0 and nums[i]<nums[0]:
                if nums[i]==target:
                    return i
                i-=1
        return -1