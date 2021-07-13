class Solution:
    def cuttingRope(self, n: int) -> int:
        if n <= 3: return n - 1
        a = n // 3
        b = n % 3
        if b == 0:
            return 3 ** a
        elif b == 1:
            return 3 ** (a - 1) * 4
        elif b == 2:
            return 3 ** a * 2
