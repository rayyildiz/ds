Data structures with [Go](http://golang.org)


### Stack 
<a href="http://en.wikipedia.org/wiki/Stack_(abstract_data_type)">Stack</a> is a data type implementation of LIFO (Last In First Out)
A stack has two important method :

* _Push_ Insert an element to stack
* _Pop_  Remove the top of the element from stack.

Here is usage of stack

    stack := NewStack()
    
    stack.Push("One")
    stack.Push("Two")

    elem, err := stack.Pop()   // elem = "Two" and err = nil



