package main

import (
	"errors"
	"fmt"
)

/*
데이터를 저장하는 선형 자료구조로, 데이터를 선입선출(FIFO, First-In-First-Out)의 순서로 처리하는 데 사용

1. Push: Adds a new item to the back of the queue.
2. Pop: Removes and returns an item from the front of the queue.
3. Front: Returns the item at the front of the queue without removing it.
4. Back: Returns the item at the back of the queue without removing it.
5. Empty: Checks if the queue is empty. Returns true if empty, false otherwise.
6. Size: Returns the number of items in the queue.
*/

type Queue struct {
	Items  []int
	Length int
}

func NewQueue() *Queue {
	return &Queue{
		Items:  []int{},
		Length: 0,
	}
}

func (q *Queue) Push(value int) {
	q.Items = append(q.Items, value)
	q.Length++
}

func (q *Queue) Pop() (int, error) {
	if q.Length == 0 {
		return 0, errors.New("Queue is empty")
	}
	item := q.Items[0]
	q.Items = q.Items[1:]
	q.Length--
	return item, nil
}

func (q *Queue) Front() (int, error) {
	if q.Length == 0 {
		return 0, errors.New("Queue is empty")
	}
	return q.Items[0], nil
}

func (q *Queue) Back() (int, error) {
	if q.Length == 0 {
		return 0, errors.New("Queue is empty")
	}
	return q.Items[q.Length-1], nil
}

func (q *Queue) IsEmpty() bool {
	return q.Length == 0
}

func (q *Queue) Size() int {
	return q.Length
}

func main() {
	q := NewQueue()
	fmt.Println(q.IsEmpty())
	q.Push(1)
	q.Push(2)
	q.Push(3)

	fmt.Println(q.Size())
	fmt.Println(q.Front())
	fmt.Println(q.Back())
	q.Pop()
	fmt.Println(q.Size())
	fmt.Println(q.Front())
	fmt.Println(q.Back())

}
