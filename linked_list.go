package goutils

import "strconv"

type ListNode struct {
	Val int
	Next *ListNode
}

type LinkedList struct {
	head *ListNode
}

func (ll *LinkedList) Head() *ListNode {
	return ll.head
}

func (ll *LinkedList) String() string {
	node := ll.head
	str := ""

	for node != nil {
		str = str + strconv.Itoa(node.Val) + "->"
		node = node.Next
	}

	str += "nil"

	return str
}

func (ll *ListNode) String() string {
	str := ""
	node := ll

	for node != nil {
		str = str + strconv.Itoa(node.Val) + "->"
		node = node.Next
	}

	str += "nil"

	return str
}

func GetLinkedListFromArray(array []int) *LinkedList {
	ll := &LinkedList{}

	if len(array) == 0 {
		return ll
	}

	root := &ListNode{Val: array[0], Next: nil}
	cur := root

	for i := 1; i < len(array); i++ {
		node := &ListNode{Val: array[i], Next: nil}
		cur.Next = node
		cur = node
	}

	ll.head = root

	return ll
}