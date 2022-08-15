/*
 * @lc app=leetcode id=2 lang=swift
 *
 * [2] Add Two Numbers
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     public var val: Int
 *     public var next: ListNode?
 *     public init() { self.val = 0; self.next = nil; }
 *     public init(_ val: Int) { self.val = val; self.next = nil; }
 *     public init(_ val: Int, _ next: ListNode?) { self.val = val; self.next = next; }
 * }
 */
class Solution {
    func(_ l1: ListNode?, _ l2: ListNode?) -> ListNode? {
        var l1 = l1
        var l2 = l2
        var head: ListNode? = nil
        var tail: ListNode? = nil
        var carry = 0
        while l1 != nil || l2 != nil {
            let sum = (l1?.val ?? 0) + (l2?.val ?? 0) + carry
            carry = sum / 10
            let node = ListNode(sum % 10)
            if head == nil {
                head = node
            } else {
                tail?.next = node
            }
            tail = node
            l1 = l1?.next
            l2 = l2?.next
        }
        if carry > 0 {
            tail?.next = ListNode(carry)
        }
        return head
    }
}
// @lc code=end

