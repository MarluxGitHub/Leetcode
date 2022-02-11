/*
 * @lc app=leetcode id=310 lang=golang
 *
 * [310] Minimum Height Trees
 */

// @lc code=start
func findMinHeightTrees(n int, edges [][]int) []int {
	// only one node, return it
	if n == 1 {
		return []int{0}
	}
	ways := make([][]int, n)
	for _, e := range edges {
		ways[e[0]] = append(ways[e[0]], e[1])
		ways[e[1]] = append(ways[e[1]], e[0])
	}
	leaves := []int{}
	wayCounter := make([]int, n)
	for i, w := range ways {
		if len(w) == 1 {
			leaves = append(leaves, i)
		}
		wayCounter[i] = len(w)
	}
	nodeLevel := make([]int, n)
	current := leaves
	level := 1
	for _, v := range current {
		nodeLevel[v] = level
	}
	for {
		level++
		next := []int{}
		for _, v := range current {
			for _, n := range ways[v] {
				wayCounter[n]--
				// more than one way, just wait
				if wayCounter[n] > 1 {
					continue
				}
				if nodeLevel[n] == 0 {
					nodeLevel[n] = level
					next = append(next, n)
				}
			}
		}
		if len(next) == 0 {
			sort.Ints(current)
			return current
		}
		current = next
	}
}
// @lc code=end

