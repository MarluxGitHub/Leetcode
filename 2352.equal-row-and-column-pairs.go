package leetcode

import "reflect"

/*
 * @lc app=leetcode id=2352 lang=golang
 *
 * [2352] Equal Row and Column Pairs
 */

// @lc code=start
func equalPairs(grid [][]int) int {
    n := len(grid)
    count := 0
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            if reflect.DeepEqual(grid[i], getColumn(grid, j)) {
                count++
            }
        }
    }
    return count
}

func getColumn(grid [][]int, j int) []int {
    n := len(grid)
    column := make([]int, n)
    for i := 0; i < n; i++ {
        column[i] = grid[i][j]
    }
    return column
}
// @lc code=end

