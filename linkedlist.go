package main

import "fmt"

// Node represents a single element in the list
type Node struct {
    value int
    next  *Node
}

// LinkedList is the list itself
type LinkedList struct {
    head *Node
}

// Insert adds a node at the end of the list
func (l *LinkedList) Insert(value int) {
    newNode := &Node{value: value}
    if l.head == nil {
        l.head = newNode
        return
    }
    current := l.head
    for current.next != nil {
        current = current.next
    }
    current.next = newNode
}

// Delete removes the first node with the given value
func (l *LinkedList) Delete(value int) {
    if l.head == nil {
        return
    }
    if l.head.value == value {
        l.head = l.head.next
        return
    }
    prev := l.head
    current := l.head.next
    for current != nil {
        if current.value == value {
            prev.next = current.next
            return
        }
        prev = current
        current = current.next
    }
}

// Find checks if a value exists in the list
func (l *LinkedList) Find(value int) bool {
    current := l.head
    for current != nil {
        if current.value == value {
            return true
        }
        current = current.next
    }
    return false
}

// Print displays all values in the list
func (l *LinkedList) Print() {
    current := l.head
    for current != nil {
        fmt.Printf("%d -> ", current.value)
        current = current.next
    }
    fmt.Println("nil")
}

// Example usage
func main() {
    list := LinkedList{}
    list.Insert(10)
    list.Insert(20)
    list.Insert(30)
    list.Print() // 10 -> 20 -> 30 -> nil

    fmt.Println("Find 20:", list.Find(20)) // true
    list.Delete(20)
    list.Print() // 10 -> 30 -> nil
}
