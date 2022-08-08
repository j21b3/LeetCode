x=2
n=10

if n == 0:
    print(1)
res = 1

if n < 0:
    x, n = 1 / x, -n

while n :
    if n & 1:
        res *= x
    x *= x
    n = n // 2

print(res)