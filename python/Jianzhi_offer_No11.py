
def minArray( numbers) -> int:
    for i in range(len(numbers)-1):
        if numbers[i]<=numbers[i+1]:
            continue
        else:
            return numbers[i+1]
    return numbers[0]

def minArray( numbers) -> int: #二分法
    low, high = 0, len(numbers) - 1
    while low < high:
        pivot = low + (high - low) // 2
        if numbers[pivot] < numbers[high]:
            high = pivot
        elif numbers[pivot] > numbers[high]:
            low = pivot + 1
        else:
            high -= 1
    return numbers[low]


print(minArray([1,3,5]))