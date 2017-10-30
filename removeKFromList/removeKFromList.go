package main

import (
	"fmt"
)

// Definition for singly-linked list:
type ListNode struct {
	Value interface{}
	Next  *ListNode
}

func removeKFromList(l *ListNode, k int) *ListNode {

	if l == nil {
		return nil
	}

	prev := new(ListNode)
	prev = nil
	ret := l

	for l != nil {
		if l.Value == k {
			if prev == nil {
				ret = l.Next
			} else {
				prev.Next = l.Next
			}
		} else {
			prev = l
		}
		l = l.Next
	}

	return ret
}

func insertKToList(l *ListNode, k int) *ListNode {

	l.Value = k
	l.Next = new(ListNode)
	l = l.Next

	return l
}

func printList(l *ListNode) {

	if l.Next == nil {
		return
	}

	fmt.Println(l.Value)
	printList(l.Next)
}

func main() {
	L := new(ListNode)
	N := new(ListNode)
	B := L

	// Case 1
	if false {
		L = insertKToList(L, 3)
		L = insertKToList(L, 1)
		L = insertKToList(L, 2)
		L = insertKToList(L, 3)
		L = insertKToList(L, 7)
		L = insertKToList(L, 1)
		L = insertKToList(L, 3)
		L = insertKToList(L, 9)
		L = insertKToList(L, 3)

		printList(B)
		N = removeKFromList(B, 3)
	}

	// Case 2
	if true {
		L = insertKToList(L, 3)
		L = insertKToList(L, 3)

		printList(B)
		N = removeKFromList(B, 3)
	}

	fmt.Println("Result:", N)
	printList(N)

}
