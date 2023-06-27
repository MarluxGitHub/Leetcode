/*
 * @lc app=leetcode id=2448 lang=golang
 *
 * [2448] Minimum Cost to Make Array Equal
 */

// @lc code=start
func minCost(nums []int, cost []int) int64 {
    n := len(nums)
    if n <= 1 {
        return 0
    }
    minCost, totalCost, maxCost := cost[0], cost[0], cost[0]
    for i := 1; i < n; i++ {
        if nums[i] > nums[i-1] {
            minCost, totalCost = cost[i], cost[i]
        } else if nums[i] < nums[i-1] {
            maxCost, totalCost = cost[i], cost[i]
        } else {
            totalCost += cost[i]
        }
        minCost = min(minCost, cost[i])
        maxCost = max(maxCost, cost[i])
    }
    if nums[0] == nums[n-1] {
		// return as int64
        return int64(min(totalCost, minCost+maxCost))
    } else {
        return int64(min(totalCost, minCost*(n-1)+maxCost))
    }
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
// @lc code=end

