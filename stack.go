package ds

import "errors"

var EmptyStack = errors.New("stack is empty")

type Stack struct {
	root *StackItem
	size int
}

type StackItem struct {
	value interface{}
	next  *StackItem
}

func NewStack() *Stack {
	stack := Stack{size: 0, root: nil}

	return &stack
}

func (stack *Stack) Push(value interface{}) {
	stack.root = &StackItem{value, stack.root}

	stack.size++
}

func (stack *Stack) Length() int {
	return stack.size
}

func (stack *Stack) Pop() (interface{}, error) {
	if stack.Length() > 0 {
		ret := stack.root
		stack.root = stack.root.next
		stack.size--

		return ret.value, nil
	}

	return nil, EmptyStack
}
