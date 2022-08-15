/*
 * @lc app=leetcode id=1202 lang=golang
 *
 * [1202] Smallest String With Swaps
 */

// @lc code=start
func smallestStringWithSwaps(s string, pairs [][]int) string {
    	n := len(s)
	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	for _, pair := range pairs {
		union(parent, pair[0], pair[1])
	}
	m := make(map[int][]byte)
	for i := range parent {
		m[find(parent, i)] = append(m[find(parent, i)], s[i])
	}
	for _, v := range m {
		sort.Slice(v, func(i, j int) bool {
			return v[i] < v[j]
		})
	}
	res := make([]byte, n)
	for i := range res {
		res[i] = m[find(parent, i)][0]
		m[find(parent, i)] = m[find(parent, i)][1:]
	}
	return string(res)
}

func find(parent []int, i int) int {
	if parent[i] != i {
		parent[i] = find(parent, parent[i])
	}
	return parent[i]
}

func union(parent []int, i, j int) {
	pi, pj := find(parent, i), find(parent, j)
	if pi != pj {
		parent[pi] = pj
	}
}
// @lc code=end

