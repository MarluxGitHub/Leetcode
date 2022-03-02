/*
 * @lc app=leetcode id=148 lang=javascript
 *
 * [148] Sort List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} head
 * @return {ListNode}
 */
 var sortList = function(head) {

    let result =[];
    let node = head;

    while(node){
        result.push(node.val);
        node = node.next;
    }

    result.sort((a,b)=> a-b)
    return helper(result)

};

function helper(array){
    return array.reduceRight((next,val)=> ({val,next}),null)
}
// @lc code=end

