package main

import "fmt"

// Node represents an element in the doubly linked list
type Node struct {
	value int
	prev  *Node
	next  *Node
}

// DoublyLinkedList manages the head and tail
type DoublyLinkedList struct {
	head *Node
	tail *Node
}

// Insert adds a new node at the end
func (l *DoublyLinkedList) Insert(value int) {
	newNode := &Node{value: value}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		return
	}
	l.tail.next = newNode
	newNode.prev = l.tail
	l.tail = newNode
}

// Delete removes the first node with the specified value
func (l *DoublyLinkedList) Delete(value int) {
	current := l.head
	for current != nil {
		if current.value == value {
			if current.prev != nil {
				current.prev.next = current.next
			} else {
				l.head = current.next
			}
			if current.next != nil {
				current.next.prev = current.prev
			} else {
				l.tail = current.prev
			}
			return
		}
		current = current.next
	}
}

// Find checks if a node with the value exists
func (l *DoublyLinkedList) Find(value int) bool {
	for node := l.head; node != nil; node = node.next {
		if node.value == value {
			return true
		}
	}
	return false
}

// PrintForward prints from head to tail
func (l *DoublyLinkedList) PrintForward() {
	fmt.Print("Forward: ")
	for node := l.head; node != nil; node = node.next {
		fmt.Printf("%d <-> ", node.value)
	}
	fmt.Println("nil")
}

// PrintBackward prints from tail to head
func (l *DoublyLinkedList) PrintBackward() {
	fmt.Print("Backward: ")
	for node := l.tail; node != nil; node = node.prev {
		fmt.Printf("%d <-> ", node.value)
	}
	fmt.Println("nil")
}

// Demo
func main() {
	list := DoublyLinkedList{}
	list.Insert(10)
	list.Insert(20)
	list.Insert(30)

	list.PrintForward()  // Forward: 10 <-> 20 <-> 30 <-> nil
	list.PrintBackward() // Backward: 30 <-> 20 <-> 10 <-> nil

	list.Delete(20)
	list.PrintForward()  // Forward: 10 <-> 30 <-> nil

	fmt.Println("Find 30:", list.Find(30)) // true
	fmt.Println("Find 99:", list.Find(99)) // false
}
