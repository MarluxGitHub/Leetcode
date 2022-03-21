/*
 * @lc app=leetcode id=4 lang=golang
 *
 * [4] Median of Two Sorted Arrays
 */

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	return getMedianOfIntArray(mergeTowIntArray(nums1, nums2))
}

func mergeTowIntArray(a, b []int) []int {
	m, n := len(a), len(b)
	c := make([]int, m+n)
	i, j, k := 0, 0, 0
	for i < m && j < n {
		if a[i] < b[j] {
			c[k] = a[i]
			i++
		} else {
			c[k] = b[j]
			j++
		}
		k++
	}
	for i < m {
		c[k] = a[i]
		i++
		k++
	}
	for j < n {
		c[k] = b[j]
		j++
		k++
	}
	return c
}

func getMedianOfIntArray(a []int) float64 {
	m := len(a)
	if m == 0 {
		return 0
	}
	if m%2 == 0 {
		return float64(a[m/2]+a[m/2-1]) / 2
	}
	return float64(a[m/2])
}
// @lc code=end

