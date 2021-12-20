/*
 * @lc app=leetcode id=25 lang=golang
 *
 * [25] Reverse Nodes in k-Group
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}

	var reversedTemp, reversedTempTail, currentSectionHead *ListNode
	cursor := 0
	result := &ListNode{}
	resultTail := result
	n := head

	for n != nil {
		if cursor == 0 {
			currentSectionHead = n
			reversedTemp = &ListNode{
				Val: n.Val,
			}
			reversedTempTail = reversedTemp
			cursor += 1
		} else if cursor < k-1 {
			reversedTemp = &ListNode{
				Val:  n.Val,
				Next: reversedTemp,
			}
			cursor += 1
		} else {
			resultTail.Next = &ListNode{
				Val:  n.Val,
				Next: reversedTemp,
			}

			resultTail = reversedTempTail
			cursor = 0
		}

		n = n.Next
		if n == nil && cursor != 0 {
			resultTail.Next = currentSectionHead
		}
	}

	return result.Next
}
// @lc code=end

