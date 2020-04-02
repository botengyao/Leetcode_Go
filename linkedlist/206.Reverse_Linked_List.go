/*
Package linkedlist ...
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
*/
package linkedlist

//ListNode val and next pointer
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var new *ListNode
	for head != nil {
		curr := head
		head = head.Next
		curr.Next = new
		new = curr
	}
	return new
}
