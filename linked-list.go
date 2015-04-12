package ds

// Node is data container of LinkedList.
type Node struct {
	next *Node
	data interface{}
}

// LinkedList container.
type LinkedList struct {
	head *Node
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
	n := Node{
		data: data,
		next: nil,
	}

	temp := list.head
	if temp == nil {
		list.head = &n
	} else {

		for {
			if temp.next == nil {
				break
			}
			temp = temp.next
		}

		temp.next = &n
	}
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
