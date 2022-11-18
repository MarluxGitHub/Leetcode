package leetcode

/*
 * @lc app=leetcode id=433 lang=golang
 *
 * [433] Minimum Genetic Mutation
 */

// @lc code=start
func minMutation(start string, end string, bank []string) int {
	if stringDiffCount(start, end) == 0 {
		return 0
	}
	bankMap := make(map[string]bool)
	for _, b := range bank {
		bankMap[b] = true
	}
	if !bankMap[end] {
		return -1
	}
	var (
		queue = []string{start}
		depth = 0
	)
	for len(queue) > 0 {
		depth++
		for i := len(queue); i > 0; i-- {
			cur := queue[0]
			queue = queue[1:]
			for k := range bankMap {
				if stringDiffCount(cur, k) == 1 {
					if stringDiffCount(k, end) == 0 {
						return depth
					}
					queue = append(queue, k)
					delete(bankMap, k)
				}
			}
		}
	}
	return -1
}

func stringDiffCount(a, b string) int {
	diff := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			diff++
		}
	}
	return diff
}
// @lc code=end

