/*
 * @lc app=leetcode id=19 lang=golang
 *
 * [19] Remove Nth Node From End of List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{
		Val:  0,
		Next: head,
	}
	for ; n > 0; n-- {
		head = head.Next
	}

	prev := dummy
	cur := dummy.Next
	for ; head != nil; head = head.Next {
		prev = prev.Next
		cur = cur.Next
	}

	prev.Next = cur.Next
	return dummy.Next
}
// @lc code=end

