/*
 * @lc app=leetcode id=147 lang=golang
 *
 * [147] Insertion Sort List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	nodes := []*ListNode{head}
	for head.Next != nil {
		head = head.Next
		nodes = append(nodes, head)
	}

	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].Val < nodes[j].Val
	})

	for i := 1; i < len(nodes); i++ {
		nodes[i-1].Next = nodes[i]
	}
	nodes[len(nodes)-1].Next = nil

	return nodes[0]
}
// @lc code=end

