package ts

// LinkedListItem wraps a value and a reference to the next LinkedListItem.
type LinkedListItem struct {
	Value interface{}
	next  *LinkedListItem
}
