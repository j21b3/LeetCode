class Solution:
    def fib(self, n: int) -> int:
        Fibo = [0, 1]
        if n == 0:
            return 0
        elif n == 1:
            return 1
        for i in range(2, n + 1):
            Fibo.append(Fibo[0] + Fibo[1])
            del Fibo[0]
        return int(Fibo[1] % 1000000007)


s = Solution()
print(s.fib(81))
