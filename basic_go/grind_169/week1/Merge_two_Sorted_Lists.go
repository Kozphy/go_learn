package week1

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(val int) *ListNode {
	return &ListNode{val, nil}
}

func print_linkList(head *ListNode) {
	var list_val = []int{}
	if head == nil {
		return
	}

	for {
		if head.Next != nil {
			list_val = append(list_val, head.Val)
			head = head.Next
		} else {
			break
		}
	}
	fmt.Println(list_val)
	// return list_val
}

func create_Linklist_Node(vals []int) *ListNode {
	var head *ListNode
	var current *ListNode
	for index, val := range vals {
		if index == 0 {
			current = NewListNode(val)
			head = current
		}
		current.Next = NewListNode(val)
		current = current.Next
	}
	return head
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	var head *ListNode = &ListNode{}
	var p1 *ListNode = head

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			p1.Next = list1
			list1 = list1.Next
		} else {
			p1.Next = list2
			list2 = list2.Next
		}
		p1 = p1.Next
	}

	if list1 != nil {
		p1.Next = list1
	}

	if list2 != nil {
		p1.Next = list2
	}
	return head.Next
}

func Execute_mergeTwoLists() {
	data_1 := []int{1, 2, 4}
	data_2 := []int{1, 3, 4}
	list1 := create_Linklist_Node(data_1)
	list2 := create_Linklist_Node(data_2)
	print_linkList(list1)
	print_linkList(list2)
	// mergeTwoLists(list1, list2)
}
