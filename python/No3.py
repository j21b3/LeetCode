s = "abcabcbb"

before, after = 0, 0
map = {}
l = -1
for after in range(len(s)):
    if s[after] not in map:
        map[s[after]] = after
    else:
        dup = map[s[after]]
        for i in range(before, dup + 1):
            map.pop(s[i])
        before = dup + 1
        map[s[after]] = after
    l = max(after - before + 1, l)

print(l)
