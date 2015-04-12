package ds

import "errors"

// ErrStack is error if stack is empty.
var ErrStack = errors.New("stack is empty")

// Stack is for root and root node.
type Stack struct {
	root *StackItem
	size int
}

// StackItem is node for data.
type StackItem struct {
	value interface{}
	next  *StackItem
}

// NewStack returns a new Stack object.
func NewStack() *Stack {
	stack := Stack{size: 0, root: nil}

	return &stack
}

// Push a item to the stack.
func (stack *Stack) Push(value interface{}) {
	stack.root = &StackItem{value, stack.root}

	stack.size++
}

// Length of the stack.
func (stack *Stack) Length() int {
	return stack.size
}

// Pop returns top element from stack.(LIFO).
func (stack *Stack) Pop() (interface{}, error) {
	if stack.Length() > 0 {
		ret := stack.root
		stack.root = stack.root.next
		stack.size--

		return ret.value, nil
	}

	return nil, ErrStack
}
