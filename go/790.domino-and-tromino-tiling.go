package leetcode

/*
 * @lc app=leetcode id=790 lang=golang
 *S
 * [790] Domino and Tromino Tiling
 */

// @lc code=start
func numTilings(n int) int {
	if (n<3) {
		return n
	}

	ppE, pE, p0, cE, c0 := 0, 1, 0, 2, 2

	for i:=3; i<=n; i++ {
		ppE = pE
		p0 = c0
		pE = cE

		cE = (ppE + p0 + pE) % 1000000007
		c0 = (ppE + ppE + p0) % 1000000007
	}

	return cE
}
// @lc code=end

