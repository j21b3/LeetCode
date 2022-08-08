class Node:
    def __init__(self, key, value):
        self.key = key
        self.value = value
        self.prev = None
        self.next = None


class LRUCache:
    def __init__(self, capacity: int):
        self.cache_page = {}
        self.capacity = capacity
        self.len = 0
        self.persudo_head = Node(-1, -1)
        self.persudo_tail = Node(-1, -1)
        self.persudo_head.next = self.persudo_tail
        self.persudo_tail.prev = self.persudo_head

    def add_head(self, node):
        node.prev = self.persudo_head
        node.next = self.persudo_head.next
        self.persudo_head.next = node
        node.next.prev = node

    def remove_node(self, node):
        node.prev.next = node.next
        node.next.prev = node.prev

    def remove_tail(self):
        node = self.persudo_tail.prev
        self.remove_node(node)
        return node

    def moveto_head(self, node):
        self.remove_node(node)
        self.add_head(node)

    def get(self, key: int) -> int:
        if key not in self.cache_page:
            return -1
        node = self.cache_page[key]
        self.moveto_head(node)
        return node.value

    def put(self, key: int, value: int) -> None:
        if key in self.cache_page:
            node = self.cache_page[key]
            node.value = value
            self.moveto_head(node)
            return None
        node = Node(key, value)
        self.add_head(node)
        self.cache_page[key] = node
        self.len += 1
        if self.len > self.capacity:
            n = self.remove_tail()
            self.cache_page.pop(n.key)
            self.len -= 1
        return None


L = LRUCache(2)
L.put(1, 1)
L.put(2, 2)
L.get(1)
L.put(3, 3)

# Your LRUCache object will be instantiated and called as such:
# obj = LRUCache(capacity)
# param_1 = obj.get(key)
# obj.put(key,value)
