package ts

import (
	"fmt"
	"testing"
)

func TestLinkedList(t *testing.T) {
	list := &LinkedList{}
	list.Push("Hello world")

	tail := list.GetTail()
	fmt.Printf("Tail: %v", tail)

	tail = "World Hello"
	fmt.Printf("Tail: %v", tail)

	_t := list.GetTail()
	fmt.Printf("Tail: %v", _t)
}
