#百度面试原题，C++可以使用逆向的单链表来解决，插入位置在头节点

class MinStack:

    def __init__(self):
        """
        initialize your data structure here.
        """
        self.sta = []
        self.len = 0
        self.minsta = []

    def push(self, x: int) -> None:
        if not self.minsta:
            self.minsta.append(x)
            self.sta.append(x)
            self.len += 1
            return
        if x <= self.minsta[-1]:
            self.minsta.append(x)
        self.sta.append(x)
        self.len += 1
        return

    def pop(self) -> None:
        if self.sta and self.minsta and self.sta[-1] == self.minsta[-1]:
            del self.sta[-1]
            del self.minsta[-1]
            self.len -= 1
        elif not self.sta:
            return
        else:
            del self.sta[-1]
            self.len -= 1

    def top(self) -> int:
        return self.sta[-1]

    def min(self) -> int:
        return self.minsta[-1]

# Your MinStack object will be instantiated and called as such:
# obj = MinStack()
# obj.push(x)
# obj.pop()
# param_3 = obj.top()
# param_4 = obj.min()