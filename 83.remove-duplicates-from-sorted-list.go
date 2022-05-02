/*
 * @lc app=leetcode id=83 lang=golang
 *
 * [83] Remove Duplicates from Sorted List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {

	c := head

	for c != nil && c.Next != nil {
		if c.Val == c.Next.Val {
			c.Next = c.Next.Next
		} else {
			c = c.Next
		}
	}

	return head
}
// @lc code=end

