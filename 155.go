package main

import(
	
)

type stack struct {
	elem int
	next *stack
}

type MinStack struct {
	top *stack
	msTop *stack
}

/** initialize your data structure here. */
func Constructor() MinStack {
	ms := MinStack{
		nil,
		nil,
	}
    return ms
}


func (this *MinStack) Push(x int)  {
    if this.top == nil&& this.msTop == nil {
		this.top = new(stack)
		this.msTop = new(stack)
		this.top.elem = x
		this.msTop.elem = x
		return
	}
	if this.msTop.elem >= x{
		temp := stack{
			x,
			this.msTop,
		}
		this.msTop = &temp
	}
	ptemp := stack{
		x,
		this.top,
	}
	this.top = &ptemp
	return
}


func (this *MinStack) Pop()  {
    if this.top.elem == this.msTop.elem {
		this.msTop = this.msTop.next
	}
	this.top = this.top.next
}


func (this *MinStack) Top() int {
    return this.top.elem
}


func (this *MinStack) GetMin() int {
    return this.msTop.elem
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */