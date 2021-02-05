package ts_test

import (
	"testing"

	"github.com/gofor-little/ts"
	"github.com/stretchr/testify/require"
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

		require.Equal(t, len(tc.elements), slice.Length())
		require.Equal(t, len(tc.elements), len(slice.GetElements()))

		for i := 0; i < len(tc.elements); i++ {
			require.Equal(t, tc.elements[i], slice.GetElement(i))
		}

		slice.Remove(tc.elements...)
		require.Equal(t, 0, slice.Length())
	}
}
