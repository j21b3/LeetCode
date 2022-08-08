class Basic:
    def __init__(self, node_num, squi_num):
        self.node_num = node_num
        self.squi_num = squi_num
        self.room = []
        self.squi_pos = []


def get_input():
    line1 = input()
    line2 = input()
    line3 = input()
    node_num, squi_num = line1.split(' ')

    basic = Basic(int(node_num), int(squi_num))

    rooms = line2.split(' ')
    for each in rooms:
        basic.room.append(int(each))

    pos = line3.split(' ')
    for each in pos:
        basic.squi_pos.append(int(each))

    return basic


def find_least_path(p1, p2):
    p_min = min(p1, p2)
    p_max = max(p1, p2)

    parent_list = [p_min]

    while p_min != 1:
        p_min = p_min // 2
        parent_list.append(p_min)

    step1 = 0
    while p_max != 1:
        p_max = p_max // 2
        step1 += 1
        if p_max in parent_list:
            break
    return parent_list.index(p_max) + step1


def run():
    data = get_input()

    # dp = [[0] * data.node_num for _ in range(data.node_num)]
    #
    # for i in range(data.node_num):
    #     dp[i][i] = 0xffffffff
    #     for j in range(i + 1, data.node_num):
    #         dp[i][j] = find_least_path(i + 1, j + 1)
    #         dp[j][i] = dp[i][j]

    step = 0
    for pos in data.squi_pos:
        ind = pos - 1
        if data.room[ind] > 0:
            data.room[ind] -= 1
            continue
        min_cost = 0xffffffff
        select_pos = 0
        for i in range(data.node_num):
            if i == ind:
                continue
            cost = find_least_path(pos, i + 1)
            if data.room[i] > 0 and min_cost > cost:
                min_cost = cost
                select_pos = i
                continue
        step += min_cost
        data.room[select_pos] -= 1
    return step


ret = run()
print(ret)
