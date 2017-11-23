package main

import (
	"fmt"
)

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := head.Next
	head.Next = nil
	for nil != p {
		t := p.Next
		p.Next = head
		head = p
		p = t
	}
	return head
}
