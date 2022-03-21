/*
 * @lc app=leetcode id=143 lang=golang
 *
 * [143] Reorder List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	var nodeList []*ListNode
	for cur := head; cur != nil; cur = cur.Next {
		nodeList = append(nodeList, cur)
	}

	i := 0
	j := len(nodeList) - 1
	nodeList[(len(nodeList)-1)/2].Next = nil
	for i < j {
		nodei := nodeList[i]
		nodej := nodeList[j]
		nodej.Next = nodei.Next
		nodei.Next = nodej
		i++
		j--
	}
}
// @lc code=end

