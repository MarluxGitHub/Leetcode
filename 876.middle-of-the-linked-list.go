/*
 * @lc app=leetcode id=876 lang=golang
 *
 * [876] Middle of the Linked List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {

	dummy := &ListNode{
		Val:  0,
		Next: head,
	}

	count := 0

	for head != nil && head.Next != nil {
		head = head.Next
		count++
	}

	if count%2 == 0 {
		count = count / 2
	} else {
		count = count / 2 + 1
	}

	head = dummy.Next

	for i := 0; i < count; i++ {
		head = head.Next
	}

	return head
}
// @lc code=end

