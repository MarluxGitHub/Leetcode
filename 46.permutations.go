/*
 * @lc app=leetcode id=46 lang=golang
 *
 * [46] Permutations
 */

// @lc code=start
func permute(nums []int) [][]int {
    	var result [][]int
	if len(nums) == 0 {
		return result
	}
	temp := []int{}
	helper(nums, temp, &result)
	return result
}

func helper(nums []int, temp []int, result *[][]int) {
	if len(temp) == len(nums) {
		r := append([]int{}, temp...)
		*result = append(*result, r)
		return
	}

	for i := 0; i < len(nums); i++ {
		if contains(temp, nums[i]) {
			continue
		}

		temp = append(temp, nums[i])
		helper(nums, temp, result)
		temp = temp[:len(temp)-1]
	}
}

func contains(nums []int, target int) bool {
	for _, n := range nums {
		if n == target {
			return true
		}
	}
	return false
}
// @lc code=end

