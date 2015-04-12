# Data structures with [Go](http://golang.org)

[![Build Status](http://img.shields.io/travis/rayyildiz/ds.svg?style=flat-square)](https://travis-ci.org/rayyildiz/ds)
[![Build status](https://ci.appveyor.com/api/projects/status/58pg386xvrekluj1?svg=true)](https://ci.appveyor.com/project/rayyildiz/ds)
[![Coverage Status](https://coveralls.io/repos/rayyildiz/ds/badge.svg?branch=master)](https://coveralls.io/r/rayyildiz/ds?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/rayyildiz/ds)](https://goreportcard.com/report/github.com/rayyildiz/ds)
[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/rayyildiz/ds)

To get this examples

    go get -u github.com/rayyildiz/ds


### Stack
[Stack](http://en.wikipedia.org/wiki/Stack_(abstract_data_type)) is a data type that  implementation of LIFO (Last In First Out)
A stack has two important method :

* _Push_ Insert an element to stack
* _Pop_  Remove the top of the element from stack.

Here is usage of stack

    stack := NewStack()

    stack.Push("One")
    stack.Push("Two")

    elem, err := stack.Pop()   // elem = "Two" and err = nil

### Queue
[Queue](http://en.wikipedia.org/wiki/Queue_(abstract_data_type)) is a data type that implementation of FIFO (First In, First Out)

A queue has two important method :

* _Enqueue_ Insert an element to queue
* _Dequeue_ Remove the first element from queue

Here is usage of stack

    q := NewQueue()

    q.Enqueue("One")
    q.Enqueue("Two")

    elem, err := q.Dequeue()   // elem = "One" and err = nil


### Linked List
Simple implementation of [linked-list](https://en.wikipedia.org/wiki/Linked_list)

Currently 2 method implemented:

* _Insert_  Insert a new element to linked list
* _Count_ Number of elements.

Here is usage of linked list

    list := NewLinkedList()

    list.Insert(3)
    list.Insert("hello")
    list.Insert(3.14)
    list.Insert(true)
