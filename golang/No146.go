package main

type Node struct {
	value, key   int
	before, next *Node
}

type LRUCache struct {
	cache      map[int]*Node
	capacity   int
	tail, head *Node
	len        int
}

func Constructor(capacity int) LRUCache {
	L := LRUCache{
		capacity: capacity,
		cache:    make(map[int]*Node, capacity),
		tail:     new(Node),
		head:     new(Node),
		len:      0,
	}
	L.tail.value = -1
	L.head.value = -1
	L.head.key = -1
	L.tail.key = -1
	L.tail.before = L.head
	L.tail.next = nil
	L.head.before = nil
	L.head.next = L.tail
	return L
}

func (this *LRUCache) RemoveNode(node *Node) {
	before, nex := node.before, node.next
	before.next = nex
	nex.before = before
}

func (this *LRUCache) InsertHead(node *Node) {
	before, nex := this.head, this.head.next
	node.before = before
	node.next = nex
	this.head.next = node
	nex.before = node
}

func (this *LRUCache) RemoveTail() *Node {
	node := this.tail.before
	this.RemoveNode(node)
	return node
}

func (this *LRUCache) MoveHead(node *Node) {
	this.RemoveNode(node)
	this.InsertHead(node)
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.cache[key]
	if !ok {
		return -1
	}
	this.MoveHead(node)
	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cache[key]; ok {
		node.value = value
		this.MoveHead(node)
		return
	}
	node := new(Node)
	node.key, node.value = key, value
	this.cache[key] = node
	this.len += 1
	this.InsertHead(node)
	if this.len > this.capacity {
		del := this.RemoveTail()
		delete(this.cache, del.key)
		this.len -= 1
	}
}
