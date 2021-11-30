package main

type MyStack struct {
	inQueue []int
	outQueue []int
}


func Constructor() MyStack {
	inQueue := []int{}
	outQueue := []int{}
	return MyStack{
		inQueue:  inQueue,
		outQueue: outQueue,
	}
}


func (this *MyStack) Push(x int)  {
	this.inQueue = append(this.inQueue, x)
}


func (this *MyStack) Pop() int {
	this.outQueue = this.inQueue[:len(this.inQueue)-1]
	popNumber := this.inQueue[len(this.inQueue)-1]
	this.inQueue = this.outQueue
	this.outQueue = []int{}
	return popNumber
}


func (this *MyStack) Top() int {
	return this.inQueue[len(this.inQueue)-1]
}


func (this *MyStack) Empty() bool {
	if len(this.inQueue) == 0 {
		return true
	} else {
		return false
	}
}
