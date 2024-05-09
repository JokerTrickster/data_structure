package main

import (
	"errors"
	"fmt"
)

// Stack is a stack data structure
/*
 스택에 대한 설명
 - 스택은 LIFO(Last In First Out)의 원리를 따르는 자료구조이다.
 - 스택은 데이터를 넣는 작업을 push, 데이터를 빼는 작업을 pop이라고 한다.
 - 스택은 배열이나 연결리스트로 구현할 수 있다.

 1. IsEmpty: Checks if the stack is empty.
2. Peek: Returns the top element of the stack without removing it.
3. Push: Adds an element to the top of the stack.
4. Pop: Removes and returns the top element of the stack.
5. Len: Returns the number of elements in the stack.

*/

type Stack struct {
	Items  []int
	Top    int
	Length int
}

func NewStack() *Stack {
	return &Stack{
		Items:  []int{},
		Top:    -1,
		Length: 0,
	}
}

func (s *Stack) IsEmpty() bool {
	return s.Length == 0
}

func (s *Stack) Peek() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("Stack is empty")
	}
	return s.Items[s.Top], nil
}

func (s *Stack) Push(value int) {
	s.Items = append(s.Items, value)
	s.Top = s.Length
	s.Length++
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("Stack is empty")
	}
	if s.Length == 1 {
		s.Items = []int{}
		s.Top = -1
		s.Length--
		return s.Items[0], nil
	}
	item := s.Items[s.Top]
	s.Items = s.Items[:s.Top]
	s.Length--
	s.Top--
	return item, nil
}

func (s *Stack) Len() int {
	return s.Length
}

func main() {
	stk := NewStack()
	stk.Push(1)
	stk.Push(2)
	stk.Push(3)
	stk.Push(4)

	fmt.Println(stk.Len())

	fmt.Println(stk.Peek())
	stk.Push(6)
	fmt.Println(stk.Peek())

	fmt.Println(stk.Len())
	stk.Pop()
	fmt.Println(stk.Len())
	fmt.Println(stk.Peek())
	stk.Push(11)
	fmt.Println(stk.Peek())
}
