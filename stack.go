package ds

import "errors"

// ErrStack is error if stack is empty.
var ErrEmptyStack = errors.New("stack is empty")

// Stack is for root and root node.
type Stack struct {
	root *stackItem
	size int
}

// StackItem is node for data.
type stackItem struct {
	value interface{}
	next  *stackItem
}

// NewStack returns a new Stack object.
func NewStack() *Stack {
	stack := Stack{size: 0, root: nil}

	return &stack
}

// Push a item to the stack.
func (stack *Stack) Push(value interface{}) {
	stack.root = &stackItem{value, stack.root}

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

	return nil, ErrEmptyStack
}
