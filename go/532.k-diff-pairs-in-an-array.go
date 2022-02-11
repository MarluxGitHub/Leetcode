package leetcode

/*
 * @lc app=leetcode id=532 lang=golang
 *
 * [532] K-diff Pairs in an Array
 */

// @lc code=start
func findPairs(nums []int, K int) int {
    if K < 0 {
        return 0
    }
    m := make(map[int]int)
    for _, v := range nums {
        m[v]++
    }
    var out int
    for k, _ := range m {
        v2 := m[k+K]
        if (K == 0 && v2 == 1) || v2 == 0 {
            continue
        }
        out++
    }
    return out
}
// @lc code=end

