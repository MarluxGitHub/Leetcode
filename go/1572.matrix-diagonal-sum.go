/*
 * @lc app=leetcode id=1572 lang=golang
 *
 * [1572] Matrix Diagonal Sum
 */

// @lc code=start
func diagonalSum(mat [][]int) int {

	n, sum := len(mat), 0

	for i := 0; i < n; i++ {
		sum += mat[i][i] + mat[i][n-i-1]
	}

	if n % 2 == 1 {
		sum -= mat[n/2][n/2]
	}

	return sum
}
// @lc code=end

