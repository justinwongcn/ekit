package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	testCases := []struct {
		name string
		in   *LinkedList[int]
		out  []int
	}{
		{
			name: "nil src",
			in:   nil,
			out:  []int{},
		},
		{
			name: "empty",
			in:   NewLinkedListOf[int](nil),
			out:  []int{},
		},
		{
			name: "single",
			in:   NewLinkedListOf([]int{1}),
			out:  []int{1},
		},
		{
			name: "multiple",
			in:   NewLinkedListOf([]int{1, 2, 3, 4, 5}),
			out:  []int{5, 4, 3, 2, 1},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := Reverse[int](tc.in)
			assert.Equal(t, tc.out, res.AsSlice())
		})
	}
}

func TestReverseSelf(t *testing.T) {
	testCases := []struct {
		name string
		in   *LinkedList[int]
		out  []int
	}{
		{
			name: "empty",
			in:   NewLinkedListOf[int](nil),
			out:  []int{},
		},
		{
			name: "single",
			in:   NewLinkedListOf([]int{42}),
			out:  []int{42},
		},
		{
			name: "multiple",
			in:   NewLinkedListOf([]int{1, 2, 3, 4, 5}),
			out:  []int{5, 4, 3, 2, 1},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ReverseSelf[int](tc.in)
			assert.Equal(t, tc.out, tc.in.AsSlice())
		})
	}
}

func TestSwapPairs(t *testing.T) {
	testCases := []struct {
		name string
		in   *LinkedList[int]
		out  []int
	}{
		{
			name: "empty",
			in:   NewLinkedListOf[int](nil),
			out:  []int{},
		},
		{
			name: "single",
			in:   NewLinkedListOf([]int{1}),
			out:  []int{1},
		},
		{
			name: "even count",
			in:   NewLinkedListOf([]int{1, 2, 3, 4}),
			out:  []int{2, 1, 4, 3},
		},
		{
			name: "odd count",
			in:   NewLinkedListOf([]int{1, 2, 3, 4, 5}),
			out:  []int{2, 1, 4, 3, 5},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			SwapPairs[int](tc.in)
			assert.Equal(t, tc.out, tc.in.AsSlice())
		})
	}
}

func TestRotateLeft(t *testing.T) {
	testCases := []struct {
		name string
		in   *LinkedList[int]
		k    int
		out  []int
	}{
		{
			name: "empty",
			in:   NewLinkedListOf[int](nil),
			k:    3,
			out:  []int{},
		},
		{
			name: "k is zero",
			in:   NewLinkedListOf([]int{1, 2, 3}),
			k:    0,
			out:  []int{1, 2, 3},
		},
		{
			name: "k less than len",
			in:   NewLinkedListOf([]int{1, 2, 3, 4, 5}),
			k:    2,
			out:  []int{3, 4, 5, 1, 2},
		},
		{
			name: "k greater than len",
			in:   NewLinkedListOf([]int{1, 2, 3, 4, 5}),
			k:    7, // 7 % 5 = 2
			out:  []int{3, 4, 5, 1, 2},
		},
		{
			name: "k equals len",
			in:   NewLinkedListOf([]int{1, 2, 3, 4}),
			k:    4,
			out:  []int{1, 2, 3, 4},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			RotateLeft[int](tc.in, tc.k)
			assert.Equal(t, tc.out, tc.in.AsSlice())
		})
	}
}

func TestRotateRight(t *testing.T) {
	testCases := []struct {
		name string
		in   *LinkedList[int]
		k    int
		out  []int
	}{
		{
			name: "empty",
			in:   NewLinkedListOf[int](nil),
			k:    3,
			out:  []int{},
		},
		{
			name: "k is zero",
			in:   NewLinkedListOf([]int{1, 2, 3}),
			k:    0,
			out:  []int{1, 2, 3},
		},
		{
			name: "k less than len",
			in:   NewLinkedListOf([]int{1, 2, 3, 4, 5}),
			k:    2,
			out:  []int{4, 5, 1, 2, 3},
		},
		{
			name: "k greater than len",
			in:   NewLinkedListOf([]int{1, 2, 3, 4, 5}),
			k:    7, // 7 % 5 = 2
			out:  []int{4, 5, 1, 2, 3},
		},
		{
			name: "k equals len",
			in:   NewLinkedListOf([]int{1, 2, 3, 4}),
			k:    4,
			out:  []int{1, 2, 3, 4},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			RotateRight[int](tc.in, tc.k)
			assert.Equal(t, tc.out, tc.in.AsSlice())
		})
	}
}
