package ds

import (
	"testing"
)

func TestNewLinkedList(t *testing.T) {
	ll := NewLinkedList()

	if ll.Count() != 0 {
		t.Error("LinkedList size should be 0")
	}
}

func TestLinkedList_Insert(t *testing.T) {
	list := NewLinkedList()

	list.Insert(3)
	list.Insert(4)
	list.Insert(5)

	if list.Count() != 3 {
		t.Error("LinkedList size should be 3")
	}
}

func TestLinkedList_Insert2(t *testing.T) {
	list := NewLinkedList()
	myVal := MyValue{"Test", 1}

	list.Insert(3)
	list.Insert("hello")
	list.Insert(3.14)
	list.Insert(true)
	list.Insert(myVal)

	if list.Count() != 5 {
		t.Error("LinkedList size should be 3")
	}
}

func BenchmarkLinkedListInsert(b *testing.B) {
	l := NewLinkedList()

	for i := 0; i < b.N; i++ {
		l.Insert(".")
	}
}

func BenchmarkLinkedListCount(b *testing.B) {
	l := NewLinkedList()
	for i := 0; i < 1000; i++ {
		l.Insert(".")
	}

	for i := 0; i < b.N; i++ {
		l.Count()
	}
}
