height = [0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1]

left_max = [0] * len(height)
right_max = [0] * len(height)

for i in range(1, len(height)):
    left_max[i] = max(left_max[i - 1], height[i - 1])

for i in range(len(height) - 2, -1, -1):
    right_max[i] = max(right_max[i + 1], height[i + 1])

print(left_max, right_max)


