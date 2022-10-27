/*
 * @lc app=leetcode id=2007 lang=golang
 *
 * [2007] Find Original Array From Doubled Array
 */

// @lc code=start
func findOriginalArray(changed []int) []int {
    if len(changed)%2 == 1 {
		return nil
	}

	m := make(map[int]int)
    for _, v := range changed {
        m[v]++
    }

    sort.Ints(changed)
    res := make([]int, 0, len(changed)/2)
    for _, v := range changed {
        if m[v] == 0 || m[2*v] == 0 {
            continue
        }
        m[v]--
        m[2*v]--
        res = append(res, v)
    }

    for _, v := range res {
        if m[2*v] != 0 {
            return nil
        }
    }

	if(len(res) == len(changed)/2) {
		return res
	}

	return nil

}
// @lc code=end

