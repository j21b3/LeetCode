class Solution:
    def cuttingRope(self, n: int) -> int:
        if n <= 3: return n - 1
        a = n // 3
        b = n % 3
        p = int(1e9+7)
        if b == 0:
            return 3 ** a % p
        elif b == 1:
            return 3 ** (a - 1) * 4 % p
        elif b == 2:
            return 3 ** a * 2 % p
