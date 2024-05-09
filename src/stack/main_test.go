package main

import (
	"errors"
	"testing"

	"github.com/go-playground/assert/v2"
)

// table driven test

func TestIsEmpty(t *testing.T) {
	tt := []struct {
		name  string
		stk   *Stack
		value []int
		want  bool
	}{
		{"stack IsEmpty Test case 1", NewStack(), []int{}, true},
		{"stack IsEmpty Test case 2", NewStack(), []int{1, 2, 3}, false},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			for _, v := range tc.value {
				tc.stk.Push(v)
			}
			if tc.stk.IsEmpty() != tc.want {
				assert.Equal(t, tc.stk.IsEmpty(), tc.want)
			}
		})
	}
}

func TestPeek(t *testing.T) {
	tt := []struct {
		name    string
		stk     *Stack
		value   []int
		want    int
		wantErr error
	}{
		{"stack Peek Test case 1", NewStack(), []int{}, 0, errors.New("Stack is empty")},
		{"stack Peek Test case 2", NewStack(), []int{1, 2, 3}, 3, nil},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			for _, v := range tc.value {
				tc.stk.Push(v)
			}
			value, err := tc.stk.Peek()
			assert.Equal(t, err, tc.wantErr)
			assert.Equal(t, value, tc.want)
		})
	}
}

func TestPush(t *testing.T) {

	tt := []struct {
		name  string
		stk   *Stack
		value []int
		want  int
	}{
		{"stack Push Test case 1", NewStack(), []int{1, 2, 3}, 3},
		{"stack Push Test case 2", NewStack(), []int{1, 2, 3, 4, 5}, 5},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			for _, v := range tc.value {
				tc.stk.Push(v)
			}
			assert.Equal(t, tc.stk.Len(), tc.want)
		})
	}
}

func TestPop(t *testing.T) {
	tt := []struct {
		name    string
		stk     *Stack
		value   []int
		want    int
		wantErr error
	}{
		{"stack Pop Test case 1", NewStack(), []int{}, 0, errors.New("Stack is empty")},
		{"stack Pop Test case 2", NewStack(), []int{1, 2, 3}, 3, nil},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			for _, v := range tc.value {
				tc.stk.Push(v)
			}
			value, err := tc.stk.Pop()
			assert.Equal(t, err, tc.wantErr)
			assert.Equal(t, value, tc.want)
		})
	}
}

func TestLen(t *testing.T) {
	tt := []struct {
		name  string
		stk   *Stack
		value []int
		want  int
	}{
		{"stack Len Test case 1", NewStack(), []int{}, 0},
		{"stack Len Test case 2", NewStack(), []int{1, 2, 3}, 3},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			for _, v := range tc.value {
				tc.stk.Push(v)
			}
			assert.Equal(t, tc.stk.Len(), tc.want)
		})
	}
}
