package reversell

// ListNode is a simple lie node object
type ListNode struct {
	Val int
	Next *ListNode
}

// ReverseList will reverse a given linkedlist
func ReverseList(head *ListNode) *ListNode {
	var prev *ListNode
	next := head

	for next != nil {
		nextT := next.Next
		next.Next = prev
		prev = next
		next = nextT
	}

	return prev
}