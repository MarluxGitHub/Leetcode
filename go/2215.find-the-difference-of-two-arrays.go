package leetcode

/*
 * @lc app=leetcode id=2215 lang=golang
 *
 * [2215] Find the Difference of Two Arrays
 */

// @lc code=start
func findDifference(nums1 []int, nums2 []int) [][]int {
	nums1 = uniqueSet(nums1)
	nums2 = uniqueSet(nums2)

	list1 := arrayDiff(nums1, nums2)
	list2 := arrayDiff(nums2, nums1)

	ret := make([][]int, 2)

	ret[0] = list1
	ret[1] = list2

	return ret
}

func uniqueSet(arr []int)[] int {
	set := make(map[int]struct{})
	for _, v := range arr {set := make(map[int]struct{})
	for _, v := range arr1 {
		set[v] = struct{}{}
	}
	for _, v := range arr2 {
		delete(set, v)
	}
	var result []int
	for k := range set {
		result = append
		set[v] = struct{}{}
	}
	var result []int
	for k := range set {
		result = append(result, k)
	}
	return result
}

func arrayDiff(arr1,arr2 []int) []int {
	set := make(map[int]struct{})
	for _, v := range arr1 {
		set[v] = struct{}{}
	}
	for _, v := range arr2 {
		delete(set, v)
	}
	var result []int
	for k := range set {
		result = append(result, k)
	}
	return result
}
// @lc code=end

