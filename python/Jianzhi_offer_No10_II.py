class Solution:
    def numWays(self, n: int) -> int:
        if n == 0:
            return 1
        elif n == 1:
            return 1
        a, b = 1, 1
        for i in range(1, n):
            a, b = b, (a + b)% 1000000007   #(a+b)mod n=(a mod n +b mod n) mod n
        return b
