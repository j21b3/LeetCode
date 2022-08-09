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
	//若是in和out都没有数据就返回-1
	if len(this.stackIn) <= 0 && len(this.stackOut) <= 0 {
		return -1
	} else {
		//要是out没有数据了就把in的倒到out里
		if len(this.stackOut) <= 0 {
			for idx := len(this.stackIn) - 1; idx >= 0; idx-- {
				this.stackOut = append(this.stackOut, this.stackIn[idx])
			}
			this.stackIn = make([]int, 0)
		}
		//从out里取数据
		ret := this.stackOut[len(this.stackOut)-1]
		this.stackOut = this.stackOut[:len(this.stackOut)-1]
		return ret
	}
}
