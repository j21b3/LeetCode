class CQueue:

    def __init__(self):
        self.stacka = []    #放入a
        self.stackb = []    #从b取，b空了就把a全倒入b

    def appendTail(self, value: int) -> None:
        self.stacka.append(value)

    def a2b(self):
        while self.stacka:
            self.stackb.append(self.stacka.pop())

    def deleteHead(self) -> int:
        if not self.stackb:
            self.a2b()
        return self.stackb.pop() if self.stackb else -1