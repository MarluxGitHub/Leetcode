package leetcode

/*
 * @lc app=leetcode id=1706 lang=golang
 *
 * [1706] Where Will the Ball Fall
 */

// @lc code=start
func findBall(grid [][]int) []int {
    m, n := len(grid), len(grid[0])
    res := make([]int, n)
    for i := 0; i < n; i++ {
        res[i] = -1
        r, c := 0, i
        for r < m {
            if grid[r][c] == 1 && c+1 < n && grid[r][c+1] == 1 {
                r, c = r+1, c+1
            } else if grid[r][c] == -1 && c-1 >= 0 && grid[r][c-1] == -1 {
                r, c = r+1, c-1
            } else {
                break
            }
        }
        if r == m {
            res[i] = c
        }
    }
    return res
}
// @lc code=end

