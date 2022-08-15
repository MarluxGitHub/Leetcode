/*
 * @lc app=leetcode id=117 lang=golang
 *
 * [117] Populating Next Right Pointers in Each Node II
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

 func connect(root *Node) *Node {
    if root == nil {
        return root
    }

    var queue []*Node

    queue = append(queue, root)

    for len(queue) > 0 {
        var lvlQueue []*Node

        for i := 0; i < len(queue); i++ {
            if queue[i].Left != nil {
                lvlQueue = append(lvlQueue, queue[i].Left)
            }

            if queue[i].Right != nil {
                lvlQueue = append(lvlQueue, queue[i].Right)
            }

            if i < len(queue)-1 {
                queue[i].Next = queue[i+1]
            }
        }

        queue = lvlQueue
    }

    return root
}

// @lc code=end

