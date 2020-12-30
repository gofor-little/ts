package ts_test

import (
	"testing"

	"github.com/gofor-little/ts"
)

func TestSlice(t *testing.T) {
	type testStruct struct {
		s string
		i int
		f float64
	}

	testCases := []struct {
		elements []interface{}
	}{
		{elements: []interface{}{"first", "second", "third"}},
		{elements: []interface{}{1, 2, 3}},
		{elements: []interface{}{1.0, 2.0, 3.0}},
		{elements: []interface{}{testStruct{s: "first", i: 1, f: 1.0}, testStruct{s: "second", i: 2, f: 2.0}, testStruct{s: "third", i: 3, f: 3.0}}},
	}

	for _, tc := range testCases {
		slice := ts.Slice{}
		slice.Add(tc.elements...)

		if slice.Length() != len(tc.elements) {
			t.Fatalf("unexpected ts.Slice length, wanted: %d, got: %d", len(tc.elements), slice.Length())
		}

		for i := 0; i < len(tc.elements); i++ {
			if slice.Elements[i] != tc.elements[i] {
				t.Fatalf("unexpected result at element: %d, wanted: %+v, got: %+v", i, tc.elements[i], slice.Elements[i])
			}
		}

		slice.Remove(tc.elements...)
		if slice.Length() != 0 {
			t.Fatalf("unexpected ts.Slice length, wanted: 0, got: %d", slice.Length())
		}
	}
}
