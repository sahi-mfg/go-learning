package main

import "fmt"

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

// Append adds a new value to the end of the list.
func (l *List[T]) Append(val T) {
	current := l
	for current.next != nil {
		current = current.next
	}
	current.next = &List[T]{val: val}
}

// PrintList traverses and prints all values in the list.
func (l *List[T]) PrintList() {
	current := l
	for current != nil {
		fmt.Print(current.val, " ")
		current = current.next
	}
	fmt.Println()
}

// Length returns the number of elements in the list.
func (l *List[T]) Length() int {
	length := 0
	current := l
	for current != nil {
		length++
		current = current.next
	}
	return length
}

func main() {
	// Create a new List that holds integers.
	l1 := &List[int]{val: 1}
	l1.Append(2)
	l1.Append(3)

	// Print the list.
	l1.PrintList() // Output: 1 2 3

	// Print the length of the list.
	fmt.Println("Length:", l1.Length()) // Output: Length: 3
}
