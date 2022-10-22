package linkedlist

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(val int) *ListNode {
	return &ListNode{val, nil}
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var prev *ListNode
	current := head
	var next *ListNode

	for {
		if current.Next != nil {
			prev = head
			head = head.Next
			current = current.Next
			next = current.Next
			current.Next = prev
			current = next
		} else {
			break
		}
	}
	return head
}

func print_reverse_List(head *ListNode) []int {
	var list_val = []int{}
	if head == nil {
		return list_val
	}

	for {
		if head.Next != nil {
			list_val = append(list_val, head.Val)
			head = head.Next
		} else {
			break
		}
	}
	return list_val
}

// TODO: To complete
func Execute_reverseList() {
	var vals []int = []int{1, 2, 3, 4, 5}
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

	ans := reverseList(head)
	print_reverse_List(ans)
}
