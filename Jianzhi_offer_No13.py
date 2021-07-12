# 剑指 Offer 13. 机器人的运动范围
ret = 0

class Solution:
    def cou(self, i, j):  # m\n在1-100之间，坐标在0-99之间
        return (i % 10) + (i // 10) + (j % 10) + (j // 10)

    def movingCount(self, m: int, n: int, k: int) -> int:
        dp = []  # 0不能去，1能去，2去过了
        global ret
        ret = 0
        for i in range(m):
            tmp = []
            for j in range(n):
                if self.cou(i, j) <= k:
                    tmp.append(1)
                else:
                    tmp.append(0)
            dp.append(tmp)
        print(dp)
        def dfs(i, j):
            if not 0 <= i < m or not 0 <= j < n or dp[i][j] != 1:
                return False
            global ret
            dp[i][j] = 2
            ret += 1  # 函数内的函数可以引用外部变量，但是不可以对其进行修改, 全局变量可以
            dfs(i + 1, j)
            dfs(i - 1, j)
            dfs(i, j + 1)
            dfs(i, j - 1)

        dfs(0, 0)
        return ret

s = Solution()
print(s.cou(15,0))
print(s.movingCount(38,15,9))