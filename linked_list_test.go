package ts_test

import (
	"testing"

	"github.com/gofor-little/ts"
)

func TestLinkedList(t *testing.T) {
	type testStruct struct {
		s string
		i int
		f float64
	}

	testCases := []struct {
		name  string
		items []interface{}
	}{
		{name: "TestLinkedList_String", items: []interface{}{"first", "second", "third"}},
		{name: "TestLinkedList_String", items: []interface{}{1, 2, 3}},
		{name: "TestLinkedList_String", items: []interface{}{testStruct{s: "first", i: 1, f: 1.0}, testStruct{s: "second", i: 2, f: 2.0}, testStruct{s: "third", i: 3, f: 3.0}}},
	}

	for _, tc := range testCases {
		list := &ts.LinkedList{}
		if !list.IsEmpty() {
			t.Fatal("ts.LinkedList should be empty")
		}

		t.Run(tc.name, func(t *testing.T) {
			for _, item := range tc.items {
				list.Push(item)
				if got := list.GetTail(); got != item {
					t.Fatalf("unexpected result returned from tc.LinkedList.GetTail(), wanted: %+v: got: %+v", item, got)
				}
			}

			for _, item := range tc.items {
				if got := list.Pop(); got != item {
					t.Fatalf("unexpected result returned from tc.LinkedList.Pop(), wanted: %+v: got: %+v", item, got)
				}
			}
		})

		if !list.IsEmpty() {
			t.Fatal("ts.LinkedList should be empty")
		}
	}
}
