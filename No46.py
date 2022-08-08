ret = []
nums = [1, 2, 3]


def trace(num, result):
    if not num:
        res = result[:]
        ret.append(res)
        result=[]
        return
    num_t = num[:]
    for each in num_t:
        result.append(each)
        num_t.remove(each)
        trace(num_t, result)
        num_t.append(each)


result = []
trace(nums, result)
print(ret)
