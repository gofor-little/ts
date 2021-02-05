package ts_test

import (
	"testing"

	"github.com/stretchr/testify/require"

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
		{name: "TestLinkedList_Int", items: []interface{}{1, 2, 3}},
		{name: "TestLinkedList_Mixed", items: []interface{}{testStruct{s: "first", i: 1, f: 1.0}, testStruct{s: "second", i: 2, f: 2.0}, testStruct{s: "third", i: 3, f: 3.0}}},
	}

	for _, tc := range testCases {
		list := &ts.LinkedList{}
		require.True(t, list.IsEmpty())

		t.Run(tc.name, func(t *testing.T) {
			for _, item := range tc.items {
				list.Push(item)
				require.Equal(t, item, list.GetTail())
			}

			for _, item := range tc.items {
				require.Equal(t, item, list.Pop())
			}
		})

		require.True(t, list.IsEmpty())
	}
}
