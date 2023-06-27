/*
 * @lc app=leetcode id=59 lang=golang
 *
 * [59] Spiral Matrix II
 */

// @lc code=start
func generateMatrix(n int) [][]int {

	if n == 0 {
		return [][]int{}
	}

	res := make([][]int, n)

	top, bottom, left, right := 0, n-1, 0, n-1

	num := 1

	// init matrix of size n with 0

	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}

	for top <= bottom && left <= right {
		for i := left; i <= right; i++ {
			res[top][i] = num
			num++
		}
		top++
		for i := top; i <= bottom; i++ {
			res[i][right] = num
			num++
		}
		right--
		if top <= bottom {
			for i := right; i >= left; i-- {
				res[bottom][i] = num
				num++
			}
			bottom--
		}
		if left <= right {
			for i := bottom; i >= top; i-- {
				res[i][left] = num
				num++
			}
			left++
		}
	}

	return res
}
// @lc code=end

