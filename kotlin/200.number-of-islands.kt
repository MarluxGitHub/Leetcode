/*
 * @lc app=leetcode id=200 lang=kotlin
 *
 * [200] Number of Islands
 */

// @lc code=start
class Solution {
    fun numIslands(grid: Array<CharArray>): Int {
        var count = 0;

        for (i in grid.indices) {
            for (j in grid[i].indices) {
                if (grid[i][j] == '1') {
                    dfs(grid, i, j)
                    count++
                }
            }
        }

        return count;
    }

    fun dfs (grid: Array<CharArray>, i: Int, j: Int) {
        if (i < 0 || i >= grid.size || j < 0 || j >= grid[i].size || grid[i][j] == '0') {
            return
        }

        grid[i][j] = '0'
        dfs(grid, i + 1, j)
        dfs(grid, i - 1, j)
        dfs(grid, i, j + 1)
        dfs(grid, i, j - 1)
    }
}
// @lc code=end

