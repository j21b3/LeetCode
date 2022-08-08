class Solution:
    def threeSum(self, nums):
        if len(nums) < 3:
            return []
        sort_num = sorted(nums)
        ans = []
        for a in range(len(sort_num)):
            if a > 0 and sort_num[a] == sort_num[a - 1]:
                continue
            tmp = -sort_num[a]
            c = len(sort_num) - 1
            for b in range(a + 1, len(sort_num)):
                if b > a + 1 and sort_num[b] == sort_num[b - 1]:
                    continue
                while b < c and sort_num[b] + sort_num[c] > tmp:
                    c -= 1
                if b == c:
                    break
                if tmp == sort_num[b] + sort_num[c]:
                    ans.append([sort_num[a], sort_num[b], sort_num[c]])
        return ans


t = [-1, 0, 1, 2, -1, -4]
s = Solution()
print(s.threeSum(t))
