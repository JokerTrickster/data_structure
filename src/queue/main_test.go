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
		q     *Queue
		value []int
		want  bool
	}{
		{"queue IsEmpty Test case 1", NewQueue(), []int{}, true},
		{"queue IsEmpty Test case 2", NewQueue(), []int{1, 2, 3}, false},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			for _, v := range tc.value {
				tc.q.Push(v)
			}
			assert.Equal(t, tc.q.IsEmpty(), tc.want)
		})
	}
}

func TestFront(t *testing.T) {
	tt := []struct {
		name    string
		q       *Queue
		value   []int
		want    int
		wantErr error
	}{
		{"queue Front Test case 1", NewQueue(), []int{}, 0, errors.New("Queue is empty")},
		{"queue Front Test case 2", NewQueue(), []int{1, 2, 3}, 1, nil},
		{"queue Front Test case 3", NewQueue(), []int{5132, 34, 3}, 5132, nil},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			for _, v := range tc.value {
				tc.q.Push(v)
			}
			value, err := tc.q.Front()
			assert.Equal(t, err, tc.wantErr)
			assert.Equal(t, value, tc.want)
		})
	}
}

func TestBack(t *testing.T) {
	tt := []struct {
		name    string
		q       *Queue
		value   []int
		want    int
		wantErr error
	}{
		{"queue Back Test case 1", NewQueue(), []int{}, 0, errors.New("Queue is empty")},
		{"queue Back Test case 2", NewQueue(), []int{1, 2, 3}, 3, nil},
		{"queue Back Test case 3", NewQueue(), []int{5132, 34, 1111}, 1111, nil},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			for _, v := range tc.value {
				tc.q.Push(v)
			}
			value, err := tc.q.Back()
			assert.Equal(t, err, tc.wantErr)
			assert.Equal(t, value, tc.want)
		})
	}
}

func TestPush(t *testing.T) {

	tt := []struct {
		name  string
		q     *Queue
		value []int
		want  int
	}{
		{"queue Push Test case 1", NewQueue(), []int{1, 2, 3}, 3},
		{"queue Push Test case 2", NewQueue(), []int{1, 2, 3, 4, 5}, 5},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			for _, v := range tc.value {
				tc.q.Push(v)
			}
			assert.Equal(t, tc.q.Size(), tc.want)
		})
	}
}

func TestPop(t *testing.T) {
	tt := []struct {
		name    string
		q       *Queue
		value   []int
		want    int
		wantErr error
	}{
		{"queue Pop Test case 1", NewQueue(), []int{}, 0, errors.New("Queue is empty")},
		{"queue Pop Test case 2", NewQueue(), []int{1, 2, 3}, 1, nil},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			for _, v := range tc.value {
				tc.q.Push(v)
			}
			value, err := tc.q.Pop()
			assert.Equal(t, err, tc.wantErr)
			assert.Equal(t, value, tc.want)
		})
	}
}

func TestSize(t *testing.T) {
	tt := []struct {
		name  string
		q     *Queue
		value []int
		want  int
	}{
		{"queue Size Test case 1", NewQueue(), []int{}, 0},
		{"queue Size Test case 2", NewQueue(), []int{1, 2, 3}, 3},
		{"queue Size Test case 3", NewQueue(), []int{5132, 34, 1111}, 3},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			for _, v := range tc.value {
				tc.q.Push(v)
			}
			assert.Equal(t, tc.q.Size(), tc.want)
		})
	}
}
