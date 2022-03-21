/*
 * @lc app=leetcode id=40 lang=golang
 *
 * [40] Combination Sum II
 */

// @lc code=start
import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	var result [][]int
	if len(candidates) == 0 {
		return result
	}
	sort.Ints(candidates)
	temp := []int{}
	helper(candidates, temp, target, &result, -1)
	return result
}

func helper(candidates []int, temp []int, target int, result *[][]int, last int) {
	if target < 0 {
		return
	}

	if target == 0 {
		r := append([]int{}, temp...)
		*result = append(*result, r)
		return
	}

	for i := last + 1; i < len(candidates); i++ {
		if i != last+1 && candidates[i] == candidates[i-1] {
			continue
		}

		t := append(temp, candidates[i])
		helper(candidates, t, target-candidates[i], result, i)
	}
}
// @lc code=end

