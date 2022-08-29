package leetcode

/*
 * @lc app=leetcode id=234 lang=golang
 *
 * [234] Palindrome Linked List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func check(node, head *ListNode) (*ListNode, bool) {
    if node.Next == nil {
        return head.Next, node.Val == head.Val
    }

    tmp, isPali := check(node.Next, head)
    return tmp.Next, isPali && tmp.Val == node.Val
}

func isPalindrome(head *ListNode) bool {
    if head == nil {
        return true
    }

    _, isPali := check(head, head)

    return isPali
}
// @lc code=end

