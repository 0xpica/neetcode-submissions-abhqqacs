/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	t1, t2 := l1, l2 
	var head,front *ListNode 
	for t1 != nil && t2 != nil {
		var temp *ListNode
		if t1.Val > t2.Val {
			temp = t2
			t2 = t2.Next
		}else {
			temp = t1
			t1 = t1.Next
		}
		if head == nil {
			head = temp
			front = head
		}else {
			head.Next = temp
			head = head.Next
		}
	}
	
	if t1 != nil {
		if head == nil {
			return t1
		}
		head.Next = t1
	}
	if t2 != nil {
		if head == nil {
			return t2
		}
		head.Next = t2
	}

	return front
}
