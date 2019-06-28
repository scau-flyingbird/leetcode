package minStack

import (
	"errors"
)

type Stack struct {
	Element []int //Element
}

func NewStack() *Stack {
	return &Stack{}
}

func (stack *Stack) Push(value ...int) {
	stack.Element = append(stack.Element, value...)
}

//返回下一个元素
func (stack *Stack) Top() (value int, err error) {
	if stack.Size() > 0 {
		return stack.Element[stack.Size()-1], nil
	} else {
		return -1, errors.New("Stack为空.")
	}
}

//返回下一个元素,并从Stack移除元素
func (stack *Stack) Pop() (err error) {
	if stack.Size() > 0 {
		stack.Element = stack.Element[:stack.Size()-1]
		return nil
	}
	return errors.New("Stack为空.") //read empty stack
}

func (stack *Stack) Size() int {
	return len(stack.Element)
}

type MinStack struct {
	MinSta *Stack
	AllSta *Stack
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		MinSta: NewStack(),
		AllSta: NewStack(),
	}
}

func (this *MinStack) Push(x int) {
	minStaTop, err := this.MinSta.Top()
	if this.MinSta.Size() <= 0 || (err == nil && x <= minStaTop) {
		this.MinSta.Push(x)
	}
	this.AllSta.Push(x)
}

func (this *MinStack) Pop() {
	allStaTop, err1 := this.AllSta.Top()
	minStaTop, err2 := this.MinSta.Top()
	this.AllSta.Pop()
	if err1 == nil && err2 == nil && allStaTop == minStaTop {
		this.MinSta.Pop()
	}
}

func (this *MinStack) Top() int {
	top, _ := this.AllSta.Top()
	return top
}

func (this *MinStack) GetMin() int {
	top, _ := this.MinSta.Top()
	return top
}
