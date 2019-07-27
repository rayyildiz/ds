package ds

import "testing"

type MyValue struct {
	Name string
	No   int
}

func TestNewStack(t *testing.T) {
	stack := NewStack()

	if stack.Length() != 0 {
		t.Error("Stack initialize error")
	}
}

func TestPush(t *testing.T) {
	myVal := MyValue{"Test", 1}

	stack := NewStack()

	stack.Push(myVal)
	if stack.Length() != 1 {
		t.Error("After push")
	}
}

func TestPop(t *testing.T) {
	stack := NewStack()

	stack.Push("One")
	stack.Push("Two")
	stack.Push("Three")

	if stack.Length() != 3 {
		t.Error("Problem with push")
	}

	elem, err := stack.Pop()

	if err != nil {
		t.Error("Error at pop")
	}

	if elem != "Three" {
		t.Error("Wrong element")
	}

	elem, err = stack.Pop()

	if err != nil {
		t.Error("Error at pop")
	}

	if elem != "Two" {
		t.Error("Wrong element")
	}

	elem, err = stack.Pop()

	if err != nil {
		t.Error("Error at pop")
	}

	if elem != "One" {
		t.Error("Wrong element")
	}

	elem, err = stack.Pop()

	if err == nil {
		t.Error("Error at pop")
	}

	if elem != nil {
		t.Error("An empty element")
	}
}

func BenchmarkStackPush(b *testing.B) {
	s := NewStack()

	for i := 0; i < b.N; i++ {
		s.Push(".")
	}
}

func BenchmarkStackCount(b *testing.B) {
	s := NewStack()

	for i := 0; i < 10; i++ {
		s.Push(".")
	}

	for i := 0; i < b.N; i++ {
		s.Length()
	}
}
