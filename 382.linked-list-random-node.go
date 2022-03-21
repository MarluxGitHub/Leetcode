package leetcode

import "math/rand"

/*
 * @lc app=leetcode id=382 lang=golang
 *
 * [382] Linked List Random Node
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type Solution struct {
    Count int
	Head *ListNode
}

func Constructor(head *ListNode) Solution {
	c := 0
	h := head

	for head != nil {
		c++
		head = head.Next
	}

	return Solution{c, h}
}


func (this *Solution) GetRandom() int {
	if this.Count == 0 {
		return 0
	}

	i := rand.Intn(this.Count)
	head := this.Head

	for i > 0 {
		head = head.Next
		i--
	}

	return head.Val
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */
// @lc code=end

