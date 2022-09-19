package main

import "container/heap"

type maxHeap []int
type minHeap []int

// 每个堆都要heap.Interface的五个方法：Len, Less, Swap, Push, Pop
// 其实只有Less的区别。

func (m maxHeap) Len() int { return len(m) }
func (m minHeap) Len() int { return len(m) }

func (m maxHeap) Less(i, j int) bool {
	return m[i] > m[j] //大顶堆
}
func (m minHeap) Less(i, j int) bool {
	return m[i] < m[j] //小顶堆
}

func (m minHeap) Swap(i, j int) { m[i], m[j] = m[j], m[i] }
func (m maxHeap) Swap(i, j int) { m[i], m[j] = m[j], m[i] }

func (m *maxHeap) Push(v interface{}) {
	*m = append(*m, v.(int))
}
func (m *minHeap) Push(v interface{}) {
	*m = append(*m, v.(int))
}

func (m *maxHeap) Pop() interface{} {
	old := *m
	n := len(old)
	v := old[n-1]
	*m = old[:n-1]
	return v
}
func (m *minHeap) Pop() interface{} {
	old := *m
	n := len(old)
	v := old[n-1]
	*m = old[:n-1]
	return v
}

type MedianFinder struct {
	maxheap *maxHeap
	minheap *minHeap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{
		new(maxHeap),
		new(minHeap),
	}
}

// AddNum 插入元素num
// 分两种情况插入：
// 1. 两个堆的大小相等，则小顶堆增加一个元素（增加的不一定是num）
// 2. 小顶堆比大顶堆多一个元素，大顶堆增加一个元素
// 这两种情况又分别对应两种情况：
// 1. num小于大顶堆的堆顶，则num插入大顶堆
// 2. num大于小顶堆的堆顶，则num插入小顶堆
// 插入完成后记得调整堆的大小使得两个堆的容量相等，或小顶堆大1
func (this *MedianFinder) AddNum(num int) {
	if this.maxheap.Len() == this.minheap.Len() {
		if this.minheap.Len() == 0 || num > (*this.minheap)[0] {
			heap.Push(this.minheap, num)
		} else {
			heap.Push(this.maxheap, num)
			heap.Push(this.minheap, heap.Pop(this.maxheap).(int))
		}
	} else {
		if num > (*this.minheap)[0] {
			heap.Push(this.minheap, num)
			heap.Push(this.maxheap, heap.Pop(this.minheap).(int))
		} else {
			heap.Push(this.maxheap, num)
		}
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.minheap.Len() == this.maxheap.Len() {
		return float64((*this.maxheap)[0])/2.0 + float64((*this.minheap)[0])/2.0
	} else {
		return float64((*this.minheap)[0])
	}
}
