
package main

import (
	"container/list"
	"fmt"
)

// Node represents a singly linked list node.
type Node struct {
	value int
	Next  *Node
}

// Print a doubly linked list.
func printList(ll *list.List) {
	for node := ll.Front(); node != nil; node = node.Next() {
		fmt.Print(node.Value, " <-> ")
	}
	fmt.Println("nil")
}

func main() {
	// Go provides a doubly linked list using container/list.
	// Unlike Java's LinkedList<T>, it stores values as any.

	// Create an empty linked list.
	ll := list.New()

	// Add elements at the end. O(1)
	ll.PushBack(12)
	ll.PushBack(11)
	ll.PushBack(14)

	fmt.Println("Length:", ll.Len())

	// Add element at the front. O(1)
	ll.PushFront(5)
	printList(ll)

	// Access first and last elements. O(1)
	fmt.Println("First:", ll.Front().Value)
	fmt.Println("Last:", ll.Back().Value)

	// Traverse forward. O(n)
	fmt.Println("Forward traversal:")
	for node := ll.Front(); node != nil; node = node.Next() {
		fmt.Println(node.Value)
	}

	// Traverse backward. O(n)
	fmt.Println("Backward traversal:")
	for node := ll.Back(); node != nil; node = node.Prev() {
		fmt.Println(node.Value)
	}

	// Insert after the first node. O(1)
	node := ll.Front()
	ll.InsertAfter(100, node)
	printList(ll)

	// Insert before the last node. O(1)
	node = ll.Back()
	ll.InsertBefore(90, node)
	printList(ll)

	// Remove last element. O(1)
	ll.Remove(ll.Back())
	printList(ll)

	// Remove first element. O(1)
	ll.Remove(ll.Front())
	printList(ll)

	// Search an element. O(n)
	target := 90
	found := false

	for node := ll.Front(); node != nil; node = node.Next() {
		if node.Value.(int) == target {
			found = true
			break
		}
	}
	fmt.Println("Found:", found)

	// Update an element. O(n) including search.
	for node := ll.Front(); node != nil; node = node.Next() {
		if node.Value.(int) == 90 {
			node.Value = 99
			break
		}
	}
	printList(ll)

	// Remove element by value. O(n)
	for node := ll.Front(); node != nil; node = node.Next() {
		if node.Value.(int) == 99 {
			ll.Remove(node)
			break
		}
	}
	printList(ll)

	// Move last node to front. O(1)
	ll.MoveToFront(ll.Back())
	printList(ll)

	// Move first node to back. O(1)
	ll.MoveToBack(ll.Front())
	printList(ll)

	// Clear the linked list.
	ll.Init()
	fmt.Println("Length after clear:", ll.Len())

	// -------- Manual Singly Linked List --------

	// Create nodes using structs and pointers.
	head := &Node{value: 10}
	second := &Node{value: 20}
	third := &Node{value: 30}

	// Connect nodes: 10 -> 20 -> 30 -> nil
	head.Next = second
	second.Next = third

	// Traverse singly linked list. O(n)
	fmt.Println("Manual linked list:")
	for curr := head; curr != nil; curr = curr.Next {
		fmt.Print(curr.value, " -> ")
	}
	fmt.Println("nil")

	// Insert at beginning. O(1)
	newNode := &Node{value: 5}
	newNode.Next = head
	head = newNode

	// Insert at end. O(n)
	lastNode := &Node{value: 40}
	curr := head

	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = lastNode

	// Print final singly linked list.
	fmt.Println("Final linked list:")
	for curr := head; curr != nil; curr = curr.Next {
		fmt.Print(curr.value, " -> ")
	}
	fmt.Println("nil")
}
