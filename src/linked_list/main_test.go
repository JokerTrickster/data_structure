package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

// test code for AddNode
// table driven test

func TestAddNode(t *testing.T) {
	list := LinkedList{}

	tt := []struct {
		name  string
		list  LinkedList
		value []int
		want  int
	}{
		{"AddNode Test case 1", list, []int{1}, 1},
		{"AddNode Test case 2", list, []int{1, 2}, 2},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			for _, v := range tc.value {
				tc.list.AddNode(v)
			}
			if tc.list.length != tc.want {
				assert.Equal(t, tc.list.length, tc.want)
			}
		})
	}
}

// test code for PrintList()
// table driven test

func TestPrintList(t *testing.T) {
	list := LinkedList{}
	tt := []struct {
		name  string
		list  LinkedList
		value []int
		want  string
	}{
		{"PrintList test case 1", list, []int{1}, "1 "},
		{"PrintList test case 2", list, []int{1, 2}, "1 2 "},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			for _, v := range tc.value {
				tc.list.AddNode(v)
			}
			assert.Equal(t, tc.list.PrintList(), tc.want)
		})
	}
}

// test code for RemoveNode()
// table driven test
func TestRemoveNode(t *testing.T) {
	list := LinkedList{}
	tt := []struct {
		name  string
		list  LinkedList
		value []int
		index int
		want  int
	}{
		{"RemoveNode test case 1", list, []int{1, 2}, 0, 1},
		{"RemoveNode test case 4", list, []int{1, 2, 3, 4, 5}, 1, 4},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			for _, v := range tc.value {
				tc.list.AddNode(v)
			}
			tc.list.RemoveNode(tc.index)
			assert.Equal(t, tc.list.length, tc.want)
		})
	}
}

// test code for GetLength()
// table driven test
func TestGetLength(t *testing.T) {
	list := LinkedList{}
	tt := []struct {
		name  string
		list  LinkedList
		value []int
		want  int
	}{
		{"GetLength Method test case 1", list, []int{1}, 1},
		{"GetLength Method test case 2", list, []int{1, 2}, 2},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			for _, v := range tc.value {
				tc.list.AddNode(v)
			}
			assert.Equal(t, tc.list.GetLength(), tc.want)
		})
	}
}

// test code for GetValue()
// table driven test

func TestGetValue(t *testing.T) {
	list := LinkedList{}
	tt := []struct {
		name  string
		list  LinkedList
		value []int
		index int
		want  int
	}{
		{"GetValue Method test case 1", list, []int{1}, 0, 1},
		{"GetValue Method test case 2", list, []int{1, 2}, 1, 2},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			for _, v := range tc.value {
				tc.list.AddNode(v)
			}
			assert.Equal(t, tc.list.GetValue(tc.index), tc.want)
		})
	}
}

// test code for InsertNode()
// table driven test

func TestInsertNode(t *testing.T) {
	list := LinkedList{}
	tt := []struct {
		name  string
		list  LinkedList
		value []int
		index int
		data  int
		want  int
	}{
		{"InsertNode Method test case 1", list, []int{1}, 0, 2, 2},
		{"InsertNode Method test case 2", list, []int{1, 2}, 1, 3, 3},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			for _, v := range tc.value {
				tc.list.AddNode(v)
			}
			tc.list.InsertNode(tc.index, tc.data)
			assert.Equal(t, tc.list.length, tc.want)
		})
	}
}
