package ds

// Node is data container of LinkedList.
type node struct {
	next *node
	data interface{}
}

// LinkedList container.
type LinkedList struct {
	head *node
}

// NewLinkedList : Create a new LinkedList object whose root is nil.
func NewLinkedList() *LinkedList {
	list := LinkedList{
		head: nil,
	}

	return &list
}

// Insert : inserting new item to linked-list.
func (list *LinkedList) Insert(data interface{}) {
	n := node{
		data: data,
		next: list.head,
	}

	list.head = &n
}

// Count : Size of elements.
func (list *LinkedList) Count() int {
	count := 0
	temp := list.head

	if temp == nil {
		return count
	}

	for {
		count++
		if temp.next == nil {
			break
		}
		temp = temp.next

	}

	return count
}
