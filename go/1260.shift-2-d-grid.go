package leetcode

/*
 * @lc app=leetcode id=1260 lang=golang
 *
 * [1260] Shift 2D Grid
 */

// @lc code=start
func shiftGrid(grid [][]int, k int) [][]int {
	n := len(grid[0])
	m := len(grid)
	k = k % (n * m)
	var buffer []int
	for i := n*m - 1; i >= 0; i-- {
		if len(buffer) < k {
			buffer = append(buffer, grid[i/n][i%n])
		}
		var index int
		if i-k >= 0 {
			index = i - k
			grid[i/n][i%n] = grid[index/n][index%n]
		} else {
			grid[i/n][i%n] = buffer[k-i-1]
		}
	}
	return grid
}
// @lc code=end

