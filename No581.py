nums=[1,2,3,3,3]

left, right = -1, -1
ma, mi = nums[0], nums[-1]

for i in range(len(nums)):
    if ma < nums[i]:
        ma = nums[i]
    else:
        right = i

for i in range(len(nums) - 1, -1, -1):
    if mi > nums[i]:
        mi = nums[i]
    else:
        left = i

print( right - left + 1 if left!= -1 else 0)