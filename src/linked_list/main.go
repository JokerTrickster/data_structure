package main

import "fmt"

// 연결 리스트 자료구조 구현
/*
	연결 리스트는 데이터 요소의 선형 집합으로, 각 요소는 다음 요소를 가리키는 방식으로 참조된다.
	연결 리스트는 배열과 달리 특정 인덱스에 접근하는데 O(n)의 시간이 걸리지 않는다.
	반면에 연결 리스트의 각 요소는 이전 요소나 다음 요소를 가리키는 참조를 가지고 있어야 하므로 저장공간의 낭비가 발생한다.
	삭제나 삽입에 대해서는 O(1)이 걸린다.

	문제 1. 연결리스트 구현
		- AddNode(data int) : 리스트에 노드를 추가
		- PrintList() : 리스트를 출력
		- RemoveNode(index int) : 리스트에서 특정 인덱스의 노드를 삭제
		- GetLength() : 리스트의 길이를 반환
		- GetValue(index int) : 리스트에서 특정 인덱스의 노드의 값을 반환
		- InsertNode(index int, data int) : 리스트에서 특정 인덱스에 노드를 추가
*/

type Node struct {
	next  *Node
	Value int
}
type LinkedList struct {
	root   *Node //첫번째 노드
	length int   //링크드 리스트 길이
}

func (ll *LinkedList) AddNode(value int) {
	if ll.root == nil {
		ll.root = &Node{Value: value}
		ll.length++
		return
	}
	newNode := &Node{Value: value}
	prev := ll.root
	for prev.next != nil {
		prev = prev.next
	}
	prev.next = newNode
	ll.length++
}
func (ll *LinkedList) PrintList() string {
	result := ""
	if ll.root == nil {
		fmt.Println("Linked List is empty")
		return ""
	}
	node := ll.root
	for node != nil {
		result += fmt.Sprintf("%d ", node.Value)
		node = node.next
	}
	return result
}

// - RemoveNode(index int) : 리스트에서 특정 인덱스의 노드를 삭제
func (ll *LinkedList) RemoveNode(index int) {
	if index < 0 || index >= ll.length {
		return
	}
	if index == 0 {
		ll.root = ll.root.next
		ll.length--
		return
	}
	node := ll.root
	//이전 노드 찾기
	for i := 0; i < index-1; i++ {
		node = node.next
	}
	node.next = node.next.next //다음 노드 줄 끊고 다다음 노드를 가리키게 함
	ll.length--
}

// - GetLength() : 리스트의 길이를 반환
func (ll *LinkedList) GetLength() int {
	return ll.length
}

// - GetIndex(index int) : 리스트에서 특정 인덱스의 노드의 값을 반환
func (ll *LinkedList) GetValue(index int) int {
	if index < 0 || index >= ll.length {
		return -1
	}
	node := ll.root
	for i := 0; i < index; i++ {
		node = node.next
	}
	return node.Value
}

// - InsertNode(index int, data int) : 리스트에서 특정 인덱스에 노드를 추가
func (ll *LinkedList) InsertNode(index int, data int) {
	if index < 0 || index > ll.length {
		return
	}
	newNode := &Node{Value: data}

	// 첫 번째 인덱스에 삽입
	if index == 0 {
		newNode.next = ll.root
		ll.root = newNode
		ll.length++
		return
	}
	node := ll.root
	for i := 0; i < index-1; i++ {
		node = node.next
	}
	newNode.next = node.next
	node.next = newNode
	ll.length++

}

func main() {
	list := LinkedList{}

	list.AddNode(1)
	list.AddNode(2)
	list.AddNode(3)
	list.AddNode(4)
	fmt.Println(list.length)

	fmt.Println("Linked List:")
	list.PrintList()

	list.RemoveNode(2)
	fmt.Println("Linked List after removing the node at index 2:")
	fmt.Println(list.length)
	list.PrintList()

	fmt.Println(list.GetValue(1))

	list.AddNode(5)
	list.AddNode(6)
	fmt.Println("Linked List after adding 5 and 6:")
	fmt.Println(list.length)
	list.PrintList()

	list.InsertNode(3, 7)
	fmt.Println("Linked List after inserting 7 at index 3:")
	fmt.Println(list.length)
	list.PrintList()
}
