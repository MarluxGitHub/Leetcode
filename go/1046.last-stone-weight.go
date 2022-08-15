/*
 * @lc app=leetcode id=1046 lang=golang
 *
 * [1046] Last Stone Weight
 */

// @lc code=start
func lastStoneWeight(stones []int) int {

	for len(stones) > 1 {
		sort.Ints(stones)
		stones[len(stones)-2] = stones[len(stones)-1] - stones[len(stones)-2]

		if(stones[len(stones)-2] == 0) {
			stones = stones[:len(stones)-2]
		} else {
			stones = stones[:len(stones)-1]
		}
	}

	ret := 0
	if len(stones) == 1 {
		ret = stones[0]
	}

	return ret
}
// @lc code=end

