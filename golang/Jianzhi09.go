package main

/*
*	两个栈倒来倒去
 */

type CQueue struct {
	stackIn  []int
	stackOut []int
}

func Constructor() CQueue {
	return CQueue{
		stackIn:  make([]int, 0),
		stackOut: make([]int, 0),
	}
}

func (this *CQueue) AppendTail(value int) {
	this.stackIn = append(this.stackIn, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.stackOut) > 0 {
		ret := this.stackOut[len(this.stackOut)-1]
		this.stackOut = this.stackOut[:len(this.stackOut)-1]
		return ret
	} else if len(this.stackIn) <= 0 {
		return -1
	} else {
		for idx := len(this.stackIn) - 1; idx >= 0; idx-- {
			this.stackOut = append(this.stackOut, this.stackIn[idx])
		}
		this.stackIn = make([]int, 0)

		ret := this.stackOut[len(this.stackOut)-1]
		this.stackOut = this.stackOut[:len(this.stackOut)-1]
		return ret
	}
}
