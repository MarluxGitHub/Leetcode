package

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

	dummy := head
	count := 0

	for head != nil && head.Next != nil {
		head = head.Next
		count++
		if(count % 2 == 0) {
			dummy = dummy.Next
		}
	}

	if(count % 2 == 1) {
		dummy = dummy.Next
	}

	return dummy
}
// @lc code=end

