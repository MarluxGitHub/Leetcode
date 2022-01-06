/*
 * @lc app=leetcode id=1094 lang=golang
 *
 * [1094] Car Pooling
 */

// @lc code=start
func carPooling(trips [][]int, capacity int) bool {
    stops := make([]int, 1001)

    for _, trip := range trips {
        stops[trip[1]] += trip[0]
        stops[trip[2]] -= trip[0]
    }

    for i := 0; capacity >= 0 && i <=1000; i++ {
        capacity -= stops[i]
    }

    return capacity >= 0
}

// @lc code=end

